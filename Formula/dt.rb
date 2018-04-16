class Dt < Formula
  desc "CLI tool for common dev tasks."
  homepage "https://github.com/chclaus/dt"
  url "https://github.com/chclaus/dt/releases/download/v0.1.2/dt_0.1.2_darwin_amd64.tar.gz"
  version "0.1.2"
  sha256 "867dec8b928200c5190175d4cb45a79f94771ad9754236d931d530f39d06bf3c"

  def install
    bin.install "dt"
  end
end
