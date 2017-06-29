require "language/go"

class Blsd < Formula
  desc "List directories in breadth-first order"
  homepage "https://github.com/junegunn/blsd"
  head "https://github.com/junegunn/blsd.git"

  depends_on "cmake" => :build
  depends_on "go" => :build
  depends_on "pkg-config" => :build

  go_resource "github.com/libgit2/git2go" do
    url "https://github.com/libgit2/git2go.git", :branch => "next"
  end

  def install
    ENV["GOPATH"] = buildpath
    Language::Go.stage_deps resources, buildpath/"src"

    cd buildpath/"src/github.com/libgit2/git2go" do
      system "git", "submodule", "update", "--init"
      system "make", "install"
    end

    mkdir_p "src/github.com/junegunn"
    ln_s buildpath, "src/github.com/junegunn/blsd"
    system "go", "build", "-ldflags", "-w", "-o", "bin/blsd"
    bin.install "bin/blsd"
  end

  test do
    system "blsd"
  end
end
