// Curve V1 AMM Swap Simulator
// Target Pool: USDC/USDT (0x4f493b7de8aac7d55f71853688b1f7c8f0243c85) StableNG

const A_PRECISION = 100n;
const PRECISION = 10n ** 18n;
const FEE_DENOMINATOR = 10n ** 10n;
const FEE = 300000n; // 0.003%
const DYNAMIC_FEE = 349949n; // 0.01%
const XP: bigint[] = [3843240489188n, 8929998181757n]; // balance change in realtime
const RATES: bigint[] = [10n ** 30n, 10n ** 30n]; 
const OFF_PEG_FEE_MULTIPLIER = 100000000000n;

function getXP(xp: bigint[], rates: bigint[]): bigint[] {
    return xp.map((x, i) => (x * rates[i]) / PRECISION);
}

function dynamicFee(xpi: bigint, xpj: bigint, fee: bigint): bigint {
    const offpegFeeMultiplier = OFF_PEG_FEE_MULTIPLIER; 
    if (offpegFeeMultiplier <= FEE_DENOMINATOR) {
        return fee;
    }

    const xps2 = (xpi + xpj) ** 2n;
    return (
        (offpegFeeMultiplier * fee) /
        (
            ((offpegFeeMultiplier - FEE_DENOMINATOR) * 4n * xpi * xpj) / xps2 +
            FEE_DENOMINATOR
        )
    );
}

function getDy(i: number, j: number, amountIn: bigint): bigint {
    const xp = getXP(XP, RATES);
    const x = xp[i] + (amountIn * RATES[i]) / PRECISION;
    const y = getY(i, j, x, xp, 2000000n);
    const dy = xp[j] - y - 1n;
    const fee = (dy * dynamicFee((xp[i] + x) / 2n, (xp[j] + y) / 2n, FEE)) / FEE_DENOMINATOR;
    return ((dy - fee) * PRECISION) / RATES[j];
}

function getD(xp: bigint[], amp: bigint): bigint {
    const n = BigInt(xp.length);
    const Ann = amp * n;
    const S = xp.reduce((acc, val) => acc + val, 0n);
    if (S === 0n) return 0n;

    let D = S;
    for (let i = 0; i < 255; i++) {
        let D_P = D;
        for (let x of xp) {
            D_P = (D_P * D) / (x * n);
        }
        const Dprev = D;
        const numerator = (Ann * S) / A_PRECISION + D_P * n;
        const denominator = ((Ann - A_PRECISION) * D) / A_PRECISION + (n + 1n) * D_P;
        D = (numerator * D) / denominator;

        if (D > Dprev) {
            if (D - Dprev <= 1n) return D;
        } else {
            if (Dprev - D <= 1n) return D;
        }
    }
    throw new Error("getD did not converge");
}

function getY(i: number, j: number, x: bigint, xp: bigint[], amp: bigint): bigint {
    const n = xp.length;
    if (i === j) throw new Error("i and j must be different");
    if (i < 0 || i >= n || j < 0 || j >= n) throw new Error("Index out of bounds");

    const N = BigInt(n);
    const D = getD(xp, amp);
    const Ann = amp * N;
    let c = D;
    let S_ = 0n;

    for (let k = 0; k < n; k++) {
        const _x = k === i ? x : k === j ? 0n : xp[k];
        if (k !== j) {
            S_ += _x;
            c = (c * D) / (_x * N);
        }
    }

    c = (c * D * A_PRECISION) / (Ann * N);
    const b = S_ + (D * A_PRECISION) / Ann;
    let y = D;

    for (let i = 0; i < 255; i++) {
        const yPrev = y;
        y = (y * y + c) / (2n * y + b - D);
        if (y === yPrev || (y > yPrev ? y - yPrev : yPrev - y) <= 1n) return y;
    }
    throw new Error("getY did not converge");
}

function main() {
    const i = 0; // Token in: USDC
    const j = 1; // Token out: USDT
    const amountIn = 1000000000000000000n; // 1e4 raw units

    try {
        const dy = getDy(i, j, amountIn);
        console.log(`Amount out (dy): ${dy.toString()}`);
    } catch (error) {
        console.error("Error calculating dy:", error);
    }
}

main();