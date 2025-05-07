const { ethers } = require("ethers");

export class MathUtils {
    static readonly N_COINS = 2n;

    static snekmateLog2(x: bigint, roundup: boolean): bigint {
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

    static cbrt(x: bigint): bigint {
        let xx = 0n;

        if (x >= 115792089237316195423570985008687907853269n * 10n ** 18n) {
            xx = x;
        } else if (x >= 115792089237316195423570985008687907853269n) {
            xx = x * 10n ** 18n;
        } else {
            xx = x * 10n ** 36n;
        }

        const log2x = this.convert(this.snekmateLog2(xx, false), "int256");
        const remainder = this.convert(log2x, "uint256") % 3n;

        let a = (((2n ** (this.convert(log2x, "uint256") / 3n)) % (2n ** 256n)) * ((1260n ** remainder) % (2n ** 256n))) / ((1000n ** remainder) % (2n ** 256n));

        a = ((2n * a) + (xx / (a * a))) / 3n;
        a = ((2n * a) + (xx / (a * a))) / 3n;
        a = ((2n * a) + (xx / (a * a))) / 3n;
        a = ((2n * a) + (xx / (a * a))) / 3n;
        a = ((2n * a) + (xx / (a * a))) / 3n;
        a = ((2n * a) + (xx / (a * a))) / 3n;
        a = ((2n * a) + (xx / (a * a))) / 3n;

        if (x >= 115792089237316195423570985008687907853269n * 10n ** 18n) {
            a = a * (10n ** 12n);
        } else if (x >= 115792089237316195423570985008687907853269n) {
            a = a * (10n ** 6n);
        }
        return a;
    }

    static max(a: bigint, b: bigint): bigint {
        return a > b ? a : b;
    }

    static min(a: bigint, b: bigint): bigint {
        return a < b ? a : b;
    }

    static abs(x: bigint): bigint {
        return x < 0n ? -x : x;
    }

    static floorDiv(a: bigint, b: bigint): bigint {
        const quotient = a / b;
        const remainder = a % b;
    
        // If remainder is not zero and signs are different, subtract 1 from quotient
        if ((remainder !== 0n) && ((a < 0n) !== (b < 0n))) {
            return quotient - 1n;
        }
        return quotient;
    }

    static sqrt(n: bigint): bigint {
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

    static convert(value: bigint, type: "int256" | "uint256"): bigint {
        const BIT_SIZE = 256n;
        const MAX_UINT = (1n << BIT_SIZE) - 1n;
        const MAX_INT = (1n << (BIT_SIZE - 1n)) - 1n;
        const MIN_INT = -(1n << (BIT_SIZE - 1n));
    
        if (type === "uint256") {
            return value & MAX_UINT;
        } else if (type === "int256") {
            const uval = value & MAX_UINT;
            return uval <= MAX_INT ? uval : uval - (1n << BIT_SIZE);
        } else {
            throw new Error("Unsupported type");
        }
    }

    static geometricMean(unsorted_x: bigint[], sort: boolean): bigint {
        let x: bigint[] = unsorted_x;
        if (sort && x[0] < x[1]) {
            x = [x[1], x[0]];
        }
        let D = x[0];
        let diff = 0n;
        for (let i = 0; i < 255; i++) {
            let D_prev = D;
            D = (D + x[0] * x[1] / D) / this.N_COINS;
            if (D > D_prev) {
                diff = D - D_prev;
            } else {
                diff = D_prev - D;
            }
            if (diff <= 1n || diff * 10n ** 18n < D) {
                return D;
            } 
        }
        throw new Error("dev: did not converge");
    }
}

export class CryptoSwap {
    static readonly N_COINS = 2n;
    private static readonly PRECISION = 10n ** 18n;
    private static readonly FEE_DENOMINATOR = 10n ** 10n;
    private static readonly A_MULTIPLIER = 10000n;
    private static readonly MIN_GAMMA = 10n ** 10n;
    private static readonly MAX_GAMMA = 2n * 10n ** 16n;
    private static readonly MIN_A = (this.N_COINS ** this.N_COINS * this.A_MULTIPLIER / 10n) ;
    private static readonly MAX_A = this.N_COINS ** this.N_COINS * this.A_MULTIPLIER * 100000n;

    private ann: bigint;
    private gamma: bigint;
    private D: bigint;
    private future_A_gamma_time: bigint;
    private price_scale: bigint;
    private mid_fee: bigint;
    private out_fee: bigint;
    private fee_gamma: bigint;
    private xp: bigint[];
    private precisions: bigint[];

    constructor(xp: bigint[], precisions: bigint[], ann: bigint, gamma: bigint, D: bigint, future_A_gamma_time: bigint, price_scale: bigint, mid_fee: bigint, out_fee: bigint, fee_gamma: bigint) {
        this.xp = xp;
        this.precisions = precisions;
        this.ann = ann;
        this.gamma = gamma;
        this.D = D;
        this.future_A_gamma_time = future_A_gamma_time;
        this.price_scale = price_scale;
        this.mid_fee = mid_fee;
        this.out_fee = out_fee;
        this.fee_gamma = fee_gamma;
    }

    public feeCalculate(xp: bigint[]): bigint {
        let f = xp[0] + xp[1];
        f = this.fee_gamma * (10n ** 18n) / (this.fee_gamma + 10n ** 18n - (10n ** 18n * MathUtils.N_COINS ** MathUtils.N_COINS) * xp[0] / f * xp[1] / f);
        return (this.mid_fee * f + this.out_fee * (10n ** 18n - f)) / (10n ** 18n);
    }

    static newtonY(ann: bigint, gamma: bigint, x: bigint[], D: bigint, i: number): bigint {
        if (ann <= this.MIN_A - 1n || ann >= this.MAX_A + 1n) {
            throw new Error("dev: unsafe values ann");
        }
        if (gamma <= this.MIN_GAMMA - 1n || gamma >= this.MAX_GAMMA + 1n) {
            throw new Error("dev: unsafe values gamma");
        }
        if (D <= 10n ** 17n - 1n || D >= 10n ** 15n * 10n ** 18n + 1n) {
            throw new Error("dev: unsafe values D");
        }

        let xj = x[1 - i];
        let y = D ** 2n / (xj * this.N_COINS ** 2n);
        let k0i = (10n ** 18n * this.N_COINS) * xj / D;

        if (k0i <= 10n ** 16n * this.N_COINS - 1n || k0i >= 10n ** 20n * this.N_COINS + 1n) {
            throw new Error("dev: unsafe values x[i]");
        }

        let convergence_limit = MathUtils.max(MathUtils.max(xj / 10n ** 14n, D / 10n ** 14n), 100n);

        for (let j = 0; j < 255; j++) {
            let y_prev = y;
            let k0 = k0i * y * this.N_COINS / D;
            let S = xj + y;
            let g1k0 = gamma + 10n ** 18n;
            if (g1k0 > k0) {
                g1k0 = g1k0 - k0 + 1n;
            } else {
                g1k0 = k0 - g1k0 + 1n;
            }
            let mul1 = 10n ** 18n * D / gamma * g1k0 / gamma * g1k0 * this.A_MULTIPLIER / ann;
            let mul2 = 10n ** 18n + (2n * 10n ** 18n) * k0 / g1k0;
            let yfprime = 10n ** 18n * y + S * mul2 + mul1;
            let dyfprime = D * mul2;
            if (yfprime < dyfprime) {
                y = y_prev / 2n;
                continue;
            } else {
                yfprime -= dyfprime
            }
            let fprime = yfprime / y;
            let y_minus = mul1 / fprime;
            let y_plus = (yfprime + 10n ** 18n * D) / fprime + y_minus * 10n ** 18n / k0;
            y_minus += 10n ** 18n * S / fprime;

            if (y_plus < y_minus) {
                y = y_prev / 2n;
            } else {
                y = y_plus - y_minus;
            }
            let diff = 0n;
            if (y > y_prev) {
                diff = y - y_prev;
            } else {
                diff = y_prev - y;
            }
            if (diff < MathUtils.max(convergence_limit, y / 10n ** 14n)) {
                let frac = y * 10n ** 18n / D;
                if (frac <= 10n ** 16n - 1n || frac >= 10n ** 20n + 1n) {
                    throw new Error("dev: unsafe values y");
                }
                return y;
            }
        }
        throw new Error("did not converge");
    }

    static newtonD(ann: bigint, gamma: bigint, x_unsorted: bigint[]): bigint {
        if (ann <= this.MIN_A - 1n || ann >= this.MAX_A + 1n) {
            throw new Error("dev: unsafe values ann");
        }
        if (gamma <= this.MIN_GAMMA - 1n || gamma >= this.MAX_GAMMA + 1n) {
            throw new Error("dev: unsafe values gamma");
        }

        let x = x_unsorted;
        if (x[0] < x[1]) {
            x = [x[1], x[0]];
        }

        if (x[0] <= 10n ** 9n - 1n || x[0] >= 10n ** 15n * 10n ** 18n + 1n) {
            throw new Error("dev: unsafe values x[0]");
        }

        if (x[1] * 10n ** 18n / x[0] > 10n ** 14n - 1n) {
            throw new Error("dev: unsafe values x[i] (input)");
        }

        let D = this.N_COINS * MathUtils.geometricMean(x, false);
        let S = x[0] + x[1];

        for (let i = 0; i < 255; i++) {
            let D_prev = D;
            let k0 = (10n ** 18n * this.N_COINS ** 2n) * x[0] / D * x[1] / D;
            let g1k0 = gamma + 10n ** 18n;
            if (g1k0 > k0) {
                g1k0 = g1k0 - k0 - 1n;
            } else {
                g1k0 = k0 - g1k0 - 1n;
            }
            let mul1 = 10n ** 18n * D / gamma * g1k0 / gamma * g1k0 * this.A_MULTIPLIER / ann;
            let mul2 = (2n * 10n ** 18n) * this.N_COINS * k0 / g1k0;
            let neg_fprime = (S + S * mul2 / 10n ** 18n) + mul1 * this.N_COINS / k0 - mul2 * D / 10n ** 18n;
            let D_plus = D * (neg_fprime + S) / neg_fprime;
            let D_minus = D * D / neg_fprime;
            if (10n ** 18n > k0) {
                D_minus += D * (mul1 / neg_fprime) / 10n ** 18n * (10n ** 18n -k0) / k0;
            } else {
                D_minus -= D * (mul1 / neg_fprime) / 10n ** 18n * (k0 - 10n ** 18n) / k0;
            }
            if (D_plus > D_minus) {
                D = D_plus - D_minus;
            } else {
                D = (D_minus - D_plus) / 2n;
            }
            let diff = 0n;
            if (D > D_prev) {
                diff = D - D_prev;
            } else {
                diff = D_prev - D;
            }
            if (diff * 10n ** 14n < MathUtils.max(10n ** 16n, D)) {
                x.forEach(_x => {
                    let frac = _x * 10n ** 18n / D;
                    if (frac <= 10n ** 16n - 1n || frac >= 10n ** 20n + 1n) {
                        throw new Error("dev: unsafe values x[i]");
                    }
                    return D;
                })
            }
        }

        throw new Error("getD did not converge");
    }

    public getDy(i: number, j: number, dx: bigint): bigint {
        let xp = this.xp;
        let D = this.D;
        if (this.future_A_gamma_time > 0n) {
            D = CryptoSwap.newtonD(this.ann, this.gamma, xp);
        }
        xp[i] += dx;
        xp = [xp[0] * this.precisions[0], xp[1] * this.price_scale * this.precisions[1] / CryptoSwap.PRECISION];
        let y = CryptoSwap.newtonY(this.ann, this.gamma, xp, D, j);
        let dy = xp[j] - y - 1n;
        xp[j] = y;
        if (j > 0) {
            dy = dy * CryptoSwap.PRECISION / this.price_scale;
        } else {
            dy /= this.precisions[0];
        }
        let fee = this.feeCalculate(xp);
        dy -= fee * dy / CryptoSwap.FEE_DENOMINATOR;
        return dy;
    }
}

async function main() {
    const provider = new ethers.providers.JsonRpcProvider("https://ethereum-rpc.publicnode.com");
    const poolAddress = "0xB576491F1E6e5E62f1d8F26062Ee822B40B0E0d4";
    const poolAbi = [
        "function balances(uint256) view returns (uint256)",
        "function price_scale() view returns (uint256)",
        "function gamma() view returns (uint256)",
        "function fee() view returns (uint256)",
        "function D() view returns (uint256)",
        "function A() view returns (uint256)",
        "function future_A_gamma_time() view returns (uint256)",
        "function mid_fee() view returns (uint256)",
        "function out_fee() view returns (uint256)",
        "function fee_gamma() view returns (uint256)",
        "function get_dy(uint256 i, uint256 j, uint256 dx) view returns (uint256)",
    ];

    const poolContract = new ethers.Contract(poolAddress, poolAbi, provider);

    async function fetchBalances(index: bigint) {
        const balance = await poolContract.balances(index);
        return BigInt(balance.toString());
    }

    async function fetchPriceScale() {
        const priceScale = await poolContract.price_scale();
        return BigInt(priceScale.toString());
    }

    async function fetchGamma() {
        const gamma = await poolContract.gamma();
        return BigInt(gamma.toString());
    }

    async function fetchD() {
        const D = await poolContract.D();
        return BigInt(D.toString());
    }

    async function fetchA() {
        const A = await poolContract.A();
        return BigInt(A.toString());
    }

    async function fetchFutureAGammaTime() {
        const futureAGammaTime = await poolContract.future_A_gamma_time();
        return BigInt(futureAGammaTime.toString());
    }

    async function fetchMidFee() {
        const midFee = await poolContract.mid_fee();
        return BigInt(midFee.toString());
    }

    async function fetchOutFee() {
        const outFee = await poolContract.out_fee();
        return BigInt(outFee.toString());
    }

    async function fetchFeeGamma() {
        const feeGamma = await poolContract.fee_gamma();
        return BigInt(feeGamma.toString());
    }

    async function fetchDy(i: number, j: number, amountIn: bigint) {
        const dy = await poolContract.get_dy(i, j, amountIn);
        return BigInt(dy.toString());
    }

    let price_scale = await fetchPriceScale();
    let A = await fetchA();
    let D = await fetchD();
    let gamma = await fetchGamma();
    let future_a_gamma_time = await fetchFutureAGammaTime();
    let mid_fee = await fetchMidFee();
    let out_fee = await fetchOutFee();
    let fee_gamma = await fetchFeeGamma();

    let i = 0; // Token in: USDC
    let j = 1; // Token out: USDT
    let amountIn = 1000000000000000000n; // 1e4 raw units

    let xp: bigint[] = [];
    xp[0] = await fetchBalances(0n);
    xp[1] = await fetchBalances(1n);

    const cryptoSwap = new CryptoSwap(xp, [1n, 1n], A, gamma, D, future_a_gamma_time, price_scale, mid_fee, out_fee, fee_gamma);

    try {
        const dy = cryptoSwap.getDy(i, j, amountIn);
        console.log(`Amount out (dy) calculated: ${dy.toString()}`);
        console.log(`Amount out (dy) from contract: ${(await fetchDy(i, j, amountIn)).toString()}`);
    } catch (error) {
        console.error("Error calculating dy:", error);
    }
}

main();