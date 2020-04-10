<h1 align="center">📈 bench</h1>

<p align="center">
  Create load profiles over a short period of time. With this tool you are able to create small
  performance profiles of different tasks in order to e.g. evaluate how migrations are affecting
  the system load.
  <br><br>
</p>


## 📦 Requirements

- Linux
- Golang(>=1.11) *(for building bench)*
- 🚧 Make *(dev dependency)*
- 🚧 staticcheck *(dev dependency)*
- 🚧 golint *(dev dependency)*


## 🔧 Installation

- Download the binary from the release page for your platform or run `go build -o bench cmd/bench/main.go`
- Copy the binary to a location in your `$PATH` (e.g. for linux `sudo cp bench /usr/local/bin/.`)


## 🚀 Usage

Simply run `bench` when you want to start and press `Ctrl-C` in order to stop and present the stats.


## 🤝 Contributing

If you are missing features or find some annoying bugs please feel free to submit an issue or a bugfix within a pull request :)


## 📝 License

© 2020 Daniel Jankowski


This project is licensed under the MIT license.


Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:


The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.


THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
