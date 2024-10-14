#!/bin/bash

function create_binding {
    contract_dir=$1
    contract=$2
    binding_dir=$3
    echo "Generating bindings for $contract"

    # Create the directory for the Go bindings
    mkdir -p "$binding_dir/${contract}"

    # Paths to the Solidity file and output directory
    file="$contract_dir/src/v3/${contract}.sol"
    dirname=$(basename "$contract" .sol)
    output_dir="out/$dirname"

    # Ensure output directory exists
    mkdir -p "$output_dir"

    # Check if the Solidity file exists
    if [ ! -f "$file" ]; then
        echo "Error: $file not found."
        return
    fi

    # Debugging output
    echo "Compiling $file to $output_dir"

    # Compile the contract using solc
    solc --bin --abi --optimize --overwrite -o "$output_dir" "$file"
    if [ $? -ne 0 ]; then
        echo "Compilation failed for $file"
        return
    fi
    echo "Compiled ${dirname}/${contract}.sol"

    # Read the ABI and bytecode
    solc_abi="$output_dir/${contract}.abi"
    solc_bin="$output_dir/${contract}.bin"

    if [ ! -f "$solc_abi" ] || [ ! -f "$solc_bin" ]; then
        echo "Error: $solc_abi or $solc_bin not found after compilation."
        return
    fi

    # Check if ABI file is valid JSON
    if ! jq . "$solc_abi" > /dev/null 2>&1; then
        echo "Error: ABI file $solc_abi is not valid JSON."
        return
    fi

    # Check if binary file exists and is not empty
    if [ ! -s "$solc_bin" ]; then
        echo "Error: Binary file $solc_bin is empty."
        return
    fi

    # Create temporary files to store ABI and bytecode
    mkdir -p data
    cp "$solc_abi" data/tmp.abi
    cp "$solc_bin" data/tmp.bin

    # Generate Go bindings using abigen
    rm -f "$binding_dir/${contract}/binding.go"
    abigen --bin=data/tmp.bin --abi=data/tmp.abi --pkg=contract${contract} --out="$binding_dir/${contract}/binding.go"

    # Clean up temporary files
    rm -rf data/tmp.abi data/tmp.bin
}

# Clean up previous bindings
rm -rf bindings/*

# List of contracts to compile and generate bindings for
avs_service_contracts="AvsServiceContract AVSTask"

for contract in $avs_service_contracts; do
    create_binding . "$contract" ./bindings
done
