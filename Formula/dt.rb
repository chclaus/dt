class Dt < Formula
  desc "CLI tool for common dev tasks."
  homepage "https://github.com/chclaus/dt"
  url "https://github.com/chclaus/dt/releases/download/v1.0.1/dt_1.0.1_darwin_amd64.tar.gz"
  version "1.0.1"
  sha256 "c4a3d96bc7d6f66e5fb8d9f8cfcaeacddb29c5e33fe91632048e4a0e5b13e620"

  def install
    bin.install "dt"
  end
end
