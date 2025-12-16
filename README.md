<div id="top"></div>

[![Contributors][contributors-shield]][contributors-url] [![Forks][forks-shield]][forks-url] [![Stargazers][stars-shield]][stars-url] [![Issues][issues-shield]][issues-url] [![MIT License][license-shield]][license-url]

<br />

<div align="center">
  <h3 align="center">Emerald</h3>

  <p align="center">
    A Proof-of-Work cryptocurrency for Minecraft economies, written in Go.
    <br />
    <a href="https://github.com/axolotlmc/emerald">Repository</a>
    ·
    <a href="https://github.com/axolotlmc/emerald/issues">Report Bug(s)</a>
  </p>
</div>

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about">About</a>
      <ul>
        <li><a href="#goals">Goals</a></li>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#running-a-node">Running a Node</a></li>
      </ul>
    </li>
    <li><a href="#why-emerald">Why Emerald?</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>

## About

Emerald is a lightweight cryptocurrency built from scratch to power in-game economies in Minecraft servers, starting with the Axolotl server. It follows a Proof-of-Work consensus model and focuses on simplicity, transparency, and educational value.

Emerald is not intended to compete with real-world financial cryptocurrencies. Instead, it serves as a practical blockchain implementation that can be tightly integrated with game mechanics, plugins, and virtual economies.

### Goals

- Provide a fair, decentralized in-game currency  
- Serve as an educational reference for blockchain concepts  
- Be simple enough to audit and extend  
- Integrate cleanly with Minecraft servers  

### Built With

- Go

<p align="right">(<a href="#top">back to top</a>)</p>

## Getting Started

Emerald can be run as a standalone node or integrated into custom tooling for Minecraft servers. Nodes communicate over a peer-to-peer network and participate in block validation via Proof-of-Work.

### Prerequisites

- Go 
- Git  
- (Optional) Minecraft server / plugin integration environment  

### Running a Node

1. Clone the repository  
   git clone https://github.com/axolotlmc/emerald.git  
   cd emerald  

2. Build the project  
   go build ./...  

3. Run the node  
   go run ./cmds/node  

Configuration options (ports, difficulty, peers, data directory) are documented in the config directory.

<p align="right">(<a href="#top">back to top</a>)</p>

## Why Emerald?

Minecraft already uses emeralds as a symbol of value and trade. This project takes that idea literally and turns it into a decentralized digital currency that players can mine, trade, and verify independently.

Emerald bridges the gap between game economies and real blockchain concepts, without real-world financial risk.

<p align="right">(<a href="#top">back to top</a>)</p>

## Roadmap

- [x] Block and header structure  
- [ ] SHA-256 based PoW  
- [ ] Basic blockchain validation  
- [ ] Peer-to-peer networking  
- [ ] Wallet implementation  
- [ ] Transaction pool (mempool)  
- [ ] Persistent storage  
- [ ] Minecraft plugin integration  
- [ ] Documentation & examples  

See the open issues page for a full list of known issues.

<p align="right">(<a href="#top">back to top</a>)</p>

## Contributing

Contributions are welcome and encouraged.

1. Fork the project  
2. Create a feature branch (git checkout -b feature/your-feature)  
3. Commit your changes (git commit -m 'Add new feature')  
4. Push to the branch (git push origin feature/your-feature)  
5. Open a Pull Request  

<p align="right">(<a href="#top">back to top</a>)</p>

## License

Distributed under the MIT License. See LICENSE for more information.

<p align="right">(<a href="#top">back to top</a>)</p>

[contributors-shield]: https://img.shields.io/github/contributors/axolotlmc/emerald?style=for-the-badge
[contributors-url]: https://github.com/axolotlmc/emerald/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/axolotlmc/emerald?style=for-the-badge
[forks-url]: https://github.com/axolotlmc/emerald/network/members
[stars-shield]: https://img.shields.io/github/stars/axolotlmc/emerald?style=for-the-badge
[stars-url]: https://github.com/axolotlmc/emerald/stargazers
[issues-shield]: https://img.shields.io/github/issues/axolotlmc/emerald?style=for-the-badge
[issues-url]: https://github.com/axolotlmc/emerald/issues
[license-shield]: https://img.shields.io/github/license/axolotlmc/emerald?style=for-the-badge
[license-url]: https://github.com/axolotlmc/emerald/blob/main/LICENSE