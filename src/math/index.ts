export function snekmateLog2(x: bigint, roundup: boolean): bigint {
    let value = x;
    let result = 0n;
    if (x >> 128n !== 0n) {
        value = x >> 128n;
        result = 128n
    }
    if (value >> 64n !== 0n) {
        value = value >> 64n;
        result += 64n
    }
    if (value >> 32n !== 0n) {
        value = value >> 32n;
        result += 32n
    }
    if (value >> 16n !== 0n) {
        value = value >> 16n;
        result += 16n
    }
    if (value >> 8n !== 0n) {   
        value = value >> 8n;
        result += 8n
    }
    if (value >> 4n !== 0n) {
        value = value >> 4n;
        result += 4n
    }
    if (value >> 2n !== 0n) {
        value = value >> 2n;
        result += 2n
    }
    if (value >> 1n !== 0n) {
        result += 1n
    }

    if (roundup && (1n << result) < x) {
        result += 1n
    }
    return result;
}

export function sqrt(n: bigint): bigint {
    if (n < 0n) throw new Error("square root of negative bigint");
    if (n < 2n) return n;

    let left = 1n;
    let right = n;
    let result = 1n;
    
    while (left <= right) {
        const mid = (left + right) / 2n;
        const square = mid * mid;
    
        if (square === n) {
        return mid;
        } else if (square < n) {
        result = mid;
        left = mid + 1n;
        } else {
        right = mid - 1n;
        }
    }
    
    return result;
}

export function cbrt(x: bigint): bigint {
    let xx = 0n;

    if (x >= 115792089237316195423570985008687907853269n * exp10(18)) {
        xx = x;
    } else if (x >= 115792089237316195423570985008687907853269n) {
        xx = x * exp10(18);
    } else {
        xx = x * exp10(36);
    }

    const log2x = convert(snekmateLog2(xx, false), "int256");
    const remainder = convert(log2x, "uint256") % 3n;

    let a = (((exp2(convert(log2x, "uint256") / 3n)) % (exp2(256))) * ((exp1260(remainder)) % (exp2(256)))) / ((exp1000(remainder)) % (exp2(256)));

    a = ((2n * a) + (xx / (a * a))) / 3n;
    a = ((2n * a) + (xx / (a * a))) / 3n;
    a = ((2n * a) + (xx / (a * a))) / 3n;
    a = ((2n * a) + (xx / (a * a))) / 3n;
    a = ((2n * a) + (xx / (a * a))) / 3n;
    a = ((2n * a) + (xx / (a * a))) / 3n;
    a = ((2n * a) + (xx / (a * a))) / 3n;

    if (x >= 115792089237316195423570985008687907853269n * exp10(18)) {
        a = a * exp10(12);
    } else if (x >= 115792089237316195423570985008687907853269n) {
        a = a * exp10(6);
    }
    return a;
}


export function geometricMean(x: bigint[], sort: boolean = false, ncoins?: bigint): bigint {
    // Case for 2 coins
    if (x.length === 2 && ncoins !== undefined) {
        let sorted = [...x];
        if (sort && sorted[0] < sorted[1]) {
            [sorted[0], sorted[1]] = [sorted[1], sorted[0]];
        }
        let D = sorted[0];
        for (let i = 0; i < 255; i++) {
            let D_prev = D;
            D = (D + sorted[0] * sorted[1] / D) / ncoins;
            let diff = D > D_prev ? D - D_prev : D_prev - D;
            if (diff <= 1n || diff * exp10(18) < D) {
                return D;
            }
        }
        throw new Error("dev: did not converge");
    }
    
    // Case for 3 coins
    if (x.length === 3) {
        let prod = (((x[0] * x[1]) / exp10(18)) * x[2]) / exp10(18);
        if (prod === 0n) {
            return 0n;
        }
        return cbrt(prod);
    }

    throw new Error("Unsupported number of coins");
}

export function convert(value: bigint, type: "int256" | "uint256"): bigint {
    const BIT_SIZE = 256n;
    const MAX_UINT = (1n << BIT_SIZE) - 1n;
    const MAX_INT = (1n << (BIT_SIZE - 1n)) - 1n;

    if (type === "uint256") {
        return value & MAX_UINT;
    } else if (type === "int256") {
        const uval = value & MAX_UINT;
        return uval <= MAX_INT ? uval : uval - (1n << BIT_SIZE);
    } else {
        throw new Error("Unsupported type");
    }
}

export function max(x: bigint, y: bigint): bigint {
    return x > y ? x : y;
}

export function min(x: bigint, y: bigint): bigint {
    return x < y ? x : y;
}

export function abs(x: bigint): bigint {
    return x < 0n ? -x : x;
}

export function exp2(n: number | bigint): bigint {
    return 2n ** BigInt(n);
}

export function exp3(n: number | bigint): bigint {
    return 3n ** BigInt(n);
}

export function exp4(n: number | bigint): bigint {
    return 4n ** BigInt(n);
}

export function exp10(n: number | bigint): bigint {
    return 10n ** BigInt(n);
}

export function exp1000(n: number | bigint): bigint {
    return 1000n ** BigInt(n);
}

export function exp1260(n: number | bigint): bigint {
    return 1260n ** BigInt(n);
}

export function pow2(n: bigint): bigint {
    if (n < 0n) throw new Error("Negative exponent");
    return n ** 2n;
}