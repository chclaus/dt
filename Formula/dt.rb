class Dt < Formula
  desc "CLI tool for common dev tasks."
  homepage "https://github.com/chclaus/dt"
  url "https://github.com/chclaus/dt/releases/download/v0.1.2/dt_0.1.2_darwin_amd64.tar.gz"
  version "0.1.2"
  sha256 "aaf5fedb3a659149b79d7fa89282e9232c298a1ede6f5ccedfa0092c189e3553"

  def install
    bin.install "dt"
  end

end
