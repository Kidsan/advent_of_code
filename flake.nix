{
  inputs = {
    naersk.url = "github:nix-community/naersk/master";
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    utils.url = "github:numtide/flake-utils";
    gomod2nix.url = "github:nix-community/gomod2nix";
  };

  outputs = { self, nixpkgs, utils, naersk, gomod2nix}:
    utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; overlays = [ gomod2nix.overlays.default]; };
        naersk-lib = pkgs.callPackage naersk { };
      in
      {
        defaultPackage = naersk-lib.buildPackage ./.;

        defaultApp = utils.lib.mkApp {
          drv = self.defaultPackage."${system}";
        };

        devShell = with pkgs; mkShell {
          buildInputs = [ cargo rustc rustfmt pre-commit rustPackages.clippy bashInteractive
            pkgs.gomod2nix
            pkg-config
            glibc
            buildPackages.go_1_19
            buildPackages.gopls
            buildPackages.gotools
            buildPackages.go-tools
            buildPackages.delve
	    nixpkgs-fmt
	    ];
          RUST_SRC_PATH = rustPlatform.rustLibSrc;
        };
      });
}
