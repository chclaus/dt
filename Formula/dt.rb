class Dt < Formula
  desc "CLI tool for common dev tasks."
  homepage "https://github.com/chclaus/dt"
  url "https://github.com/chclaus/dt/releases/download/v1.0.4/dt_1.0.4_darwin_amd64.tar.gz"
  version "1.0.4"
  sha256 "8acf7dc276a8b74977df104ae855eb4bbbdd9474c0f23708cc5ad0d70ff924da"

  def install
    bin.install "dt"
  end
end
