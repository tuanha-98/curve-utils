const { ethers } = require("ethers");

export class StableSwap {
    private static readonly A_PRECISION = 100n;
    private static readonly PRECISION = 10n ** 18n;
    private static readonly FEE_DENOMINATOR = 10n ** 10n;

    private xp: bigint[];
    private rates: bigint[];
    private FEE: bigint;
    private OFF_PEG_FEE_MULTIPLIER: bigint;

    constructor(xp: bigint[], rates: bigint[], fee: bigint, offpegFeeMultiplier: bigint) {
        this.xp = xp;
        this.rates = rates;
        this.FEE = fee;
        this.OFF_PEG_FEE_MULTIPLIER = offpegFeeMultiplier;
    }

    private getXP(): bigint[] {
        return this.xp.map((x, i) => (x * this.rates[i]) / StableSwap.PRECISION);
    }

    private dynamicFee(xpi: bigint, xpj: bigint, fee: bigint): bigint {
        const offpegFeeMultiplier = this.OFF_PEG_FEE_MULTIPLIER;
        if (offpegFeeMultiplier <= StableSwap.FEE_DENOMINATOR) {
            return fee;
        }

        const xps2 = (xpi + xpj) ** 2n;
        return (
            (offpegFeeMultiplier * fee) /
            (
                ((offpegFeeMultiplier - StableSwap.FEE_DENOMINATOR) * 4n * xpi * xpj) / xps2 +
                StableSwap.FEE_DENOMINATOR
            )
        );
    }

    private getD(xp: bigint[], amp: bigint): bigint {
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
            const numerator = (Ann * S) / StableSwap.A_PRECISION + D_P * n;
            const denominator = ((Ann - StableSwap.A_PRECISION) * D) / StableSwap.A_PRECISION + (n + 1n) * D_P;
            D = (numerator * D) / denominator;

            if (D > Dprev) {
                if (D - Dprev <= 1n) return D;
            } else {
                if (Dprev - D <= 1n) return D;
            }
        }
        throw new Error("getD did not converge");
    }

    private getY(i: number, j: number, x: bigint, xp: bigint[], amp: bigint): bigint {
        const n = xp.length;
        if (i === j) throw new Error("i and j must be different");
        if (i < 0 || i >= n || j < 0 || j >= n) throw new Error("Index out of bounds");

        const N = BigInt(n);
        const D = this.getD(xp, amp);
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

        c = (c * D * StableSwap.A_PRECISION) / (Ann * N);
        const b = S_ + (D * StableSwap.A_PRECISION) / Ann;
        let y = D;

        for (let i = 0; i < 255; i++) {
            const yPrev = y;
            y = (y * y + c) / (2n * y + b - D);
            if (y === yPrev || (y > yPrev ? y - yPrev : yPrev - y) <= 1n) return y;
        }
        throw new Error("getY did not converge");
    }

    public getDy(i: number, j: number, amountIn: bigint): bigint {
        const xp = this.getXP();
        const x = xp[i] + (amountIn * this.rates[i]) / StableSwap.PRECISION;
        const y = this.getY(i, j, x, xp, 2000000n);
        const dy = xp[j] - y - 1n;
        const fee = (dy * this.dynamicFee((xp[i] + x) / 2n, (xp[j] + y) / 2n, this.FEE)) / StableSwap.FEE_DENOMINATOR;
        return ((dy - fee) * StableSwap.PRECISION) / this.rates[j];
    }
}

// Example usage
async function main() {
    const provider = new ethers.providers.JsonRpcProvider("https://ethereum-rpc.publicnode.com");
    const poolAddress = "0x4f493B7dE8aAC7d55F71853688b1F7C8F0243C85";
    const poolAbi = [
        "function get_balances() view returns (uint256[])",
        "function stored_rates() view returns (uint256[])",
        "function get_dy(int128 i, int128 j, uint256 dx) view returns (uint256)",
        "function fee() view returns (uint256)",
        "function offpeg_fee_multiplier() view returns (uint256)",
    ];

    const poolContract = new ethers.Contract(poolAddress, poolAbi, provider);

    async function fetchBalances() {
        const balances = await poolContract.get_balances();
        return balances.map((balance: any) => BigInt(balance.toString()));
    }

    async function fetchRates() {
        const rates = await poolContract.stored_rates();
        return rates.map((rate: any) => BigInt(rate.toString()));
    }

    async function fetchDy(i: number, j: number, amountIn: bigint) {
        const dy = await poolContract.get_dy(i, j, amountIn);
        return BigInt(dy.toString());
    }

    async function fetchFee() {
        const fee = await poolContract.fee();
        return BigInt(fee.toString());
    }

    async function fetchOffpegFeeMultiplier() {
        const offpegFeeMultiplier = await poolContract.offpeg_fee_multiplier();
        return BigInt(offpegFeeMultiplier.toString());
    }

    const xp = await fetchBalances();
    const rates = await fetchRates();
    const fee = await fetchFee();
    const offpegFeeMultiplier = await fetchOffpegFeeMultiplier();
    const stableSwap = new StableSwap(xp, rates, fee, offpegFeeMultiplier);

    const i = 0; // Token in: USDC
    const j = 1; // Token out: USDT
    const amountIn = 1000000000000000000n; // 1e4 raw units

    try {
        const dy = stableSwap.getDy(i, j, amountIn);
        console.log(`Amount out (dy) calculated: ${dy.toString()}`);
        console.log(`Amount out (dy) from contract: ${(await fetchDy(i, j, amountIn)).toString()}`);
    } catch (error) {
        console.error("Error calculating dy:", error);
    }
}

main();