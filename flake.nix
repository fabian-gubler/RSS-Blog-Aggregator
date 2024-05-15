{
  description = "RSSFlow Development Environment";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-23.11";
    nixpkgs-unstable.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs = { nixpkgs, nixpkgs-unstable, flake-utils, ... }:

    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
        unstable = import nixpkgs-unstable {
          inherit system;
        };
      in
      {
        devShell = pkgs.mkShell {
          name = "dev";
          buildInputs = [
            pkgs.sqlc
            pkgs.goose
            unstable.go
          ];

        };
      }
    );
}
