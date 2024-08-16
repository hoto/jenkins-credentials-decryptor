{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/release-24.05";
    flake-parts.url = "github:hercules-ci/flake-parts";
    flake-root.url = "github:srid/flake-root";
    gomod2nix.url = "github:nix-community/gomod2nix";
    git-hooks = {
      url = "github:cachix/git-hooks.nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    systems.url = "github:nix-systems/default";
    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = inputs:
    inputs.flake-parts.lib.mkFlake {inherit inputs;} {
      systems = inputs.nixpkgs.lib.trivial.id (import inputs.systems);
      imports = [
        inputs.flake-root.flakeModule
        inputs.git-hooks.flakeModule
        inputs.treefmt-nix.flakeModule
      ];
      perSystem = {
        config,
        self',
        inputs',
        system,
        ...
      }: let
        pkgs = import inputs.nixpkgs {
          inherit system;
          overlays = [inputs.gomod2nix.overlays.default (final: prev: {})];
          config = {};
        };
        version = inputs.self.shortRev or "development";
        jenkins-credentials-decryptor = pkgs.buildGoApplication {
          inherit version;
          pname = "jenkins-credentials-decryptor";
          src = ./.;
          ldflags = ["-X github.com/lafrenierejm/gron/cmd.Version=${version}"];
          modules = ./gomod2nix.toml;
          meta = with pkgs.lib; {
            description = "Command line tool for decrypting and dumping Jenkins credentials";
            homepage = "https://github.com/hoto/jenkins-credentials-decryptor";
            license = licenses.mit;
            maintainers = with maintainers; [lafrenierejm];
          };
        };
      in {
        # Per-system attributes can be defined here. The self' and inputs'
        # module parameters provide easy access to attributes of the same
        # system.
        packages = {
          inherit jenkins-credentials-decryptor;
          default = jenkins-credentials-decryptor;
        };

        apps = {
          inherit jenkins-credentials-decryptor;
          default = jenkins-credentials-decryptor;
        };

        # Auto formatters.
        treefmt.config = {
          projectRootFile = ".git/config";
          package = pkgs.treefmt;
          flakeCheck = false; # use pre-commit's check instead
          programs = {
            alejandra.enable = true;
          };
        };

        pre-commit = {
          check.enable = true;
          settings.hooks = {
            editorconfig-checker.enable = true;
            treefmt.enable = true;
            typos.enable = true;
          };
        };

        devShells.default = pkgs.mkShell {
          # Inherit all of the pre-commit hooks.
          inputsFrom = [config.pre-commit.devShell];
          packages =
            config.pre-commit.settings.enabledPackages
            ++ (with pkgs; [
              go-tools
              godef
              gomod2nix
              gopls
              gotools
              (mkGoEnv {pwd = ./.;})
            ]);
        };
      };
    };
}
