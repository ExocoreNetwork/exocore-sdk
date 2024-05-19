const hexString = "0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000201ca41d12637f4a7096d60506199f47f21fe8a9a2dde94327671a7d9fd4589dc0";

const hexData = hexString.slice(2);

const byteArray = [];
for (let i = 0; i < hexData.length; i += 2) {
    const hexByte = hexData.substr(i, 2);
    byteArray.push(parseInt(hexByte, 16));
}

while (byteArray.length < 32) {
    byteArray.unshift(0);
}

console.log(byteArray.slice(0, 32));