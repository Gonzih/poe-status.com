with import <nixpkgs> {};
pkgs.mkShell {
  buildInputs = [
    glib pkgconfig
    rustup gnumake
    protobuf clojure
    go_1_13
    nmap
  ];
}
