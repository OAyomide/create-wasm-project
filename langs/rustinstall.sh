#!/bin/bash

function install_curl {
    curlv=`curl -V`;
    if [ -z "$curlv" ];
    then
        echo "cURL not installed...Installing cURL first"
        sudo apt update && sudo apt upgrade
        sudo apt install curl -y
        echo "##########################"
        echo "cURL has been installed..."
    else
        echo "cURL has already been installed"
    fi
}

function install_wasmpack {
    install_curl
    wasmpackv=`wasm-pack -V`;
    
    if [ -z "$wasmpackv" ];
    then
        echo "Rust has been installed but wasm-pack has not. Proceeding to install wasmpack"
        curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh
        echo " "
        echo "################################################"
        echo "wasm-pack has been installed.."
    else
        echo ""
        echo "################################################"
        echo "wasm-pack has been installed.."
    fi
}

function install_rust {
    rustv=`rustc -V`;
    if [ -z "$rustv" ];
    then
        # here is where we install rust
        echo "Rust has not been installed... "
        install_curl
        echo "cURL already installed, proceeding to install rustc"
        echo "################################################"
        echo " "
        curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
        echo " "
        echo "################################################"
        echo "Successfully installed Rust"
    else
        echo "Rust has been installed...skipping"
    fi
    
    install_wasmpack
    echo ""
    echo "Seems you're ready to go. Everything needed has been installed"
}

install_rust