class Dt < Formula
  desc "CLI tool for common dev tasks."
  homepage "https://github.com/chclaus/dt"
  url "https://github.com/chclaus/dt/releases/download/v1.0.3/dt_1.0.3_darwin_amd64.tar.gz"
  version "1.0.3"
  sha256 "ffc88837c3d247293c34ef8f3c4db5904afee8631513afa53b211d19cfd940c1"

  def install
    bin.install "dt"
  end
end
