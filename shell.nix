with import <nixpkgs> {};
pkgs.mkShell {
  buildInputs = [
    glib pkgconfig
    rustup gnumake
    protobuf clojure
    nmap
  ];
}
