const { ethers } = require("ethers");

export class MathUtils {
    static readonly N_COINS = 2n;
    private static readonly A_MULTIPLIER = 10000n;
    private static readonly MIN_GAMMA = 10n ** 10n;
    private static readonly MAX_GAMMA_SMALL = 2n * 10n ** 16n;
    private static readonly MAX_GAMMA = 199n * 10n ** 15n;
    private static readonly MIN_A = (this.N_COINS ** this.N_COINS * this.A_MULTIPLIER / 10n) ;
    private static readonly MAX_A = this.N_COINS ** this.N_COINS * this.A_MULTIPLIER * 1000n;

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

    private static _newtonY(ann: bigint, gamma: bigint, x: bigint[], D: bigint, i: number, lim_mul: bigint): bigint {
        const x_j = x[1 - i];
        let y = D ** 2n / (x_j * this.N_COINS ** 2n)
        const k0_i = (10n ** 18n * this.N_COINS) * x_j / D; // 

        if ((k0_i < (10n ** 36n / lim_mul)) || (k0_i > lim_mul)) {
            throw new Error("dev: unsafe values x[i]");
        }
        
        const convergence_limit = this.max(this.max((x_j / 10n ** 14n), (D / 10n ** 14n)), 100n);

        for (let j = 0; j < 255; ++j) {
            const y_prev = y;
            
            const k0 = k0_i * y * this.N_COINS / D;
            let S = x_j + y;

            let g1k0 = gamma + 10n ** 18n;
            if (g1k0 > k0) {
                g1k0 = g1k0 - k0 + 1n;
            } else {
                g1k0 = k0 - g1k0 + 1n;
            }

            const mul1 = (((((10n ** 18n * D) / gamma) * g1k0) / gamma) * g1k0 * 10000n) / ann;
            const mul2 = 10n ** 18n + (2n * 10n ** 18n) * k0 / g1k0;

            let yfprime = 10n ** 18n * y + S * mul2 + mul1;
            const dyfprime = D * mul2;

            if (yfprime < dyfprime) {
                y = y_prev / 2n;
                continue;
            } else {
                yfprime -= dyfprime;
            }
            const fprime = yfprime / y;
            let y_minus = mul1 / fprime;
            const y_plus = (yfprime + 10n ** 18n * D) / fprime + y_minus * 10n ** 18n / k0;
            y_minus += 10n ** 18n * (S / fprime);



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
            if (diff < this.max(convergence_limit, y / 10n ** 14n)) {
                return y;
            }
        }
        throw new Error("dev: did not converge");
    }

    public static newtonY(ann: bigint, gamma: bigint, x: bigint[], D: bigint, i: number): bigint {
        if (ann <= this.MIN_A - 1n || ann >= this.MAX_A + 1n) {
            throw new Error("dev: unsafe values ann");
        }
        if (gamma <= this.MIN_GAMMA - 1n || gamma >= this.MAX_GAMMA + 1n) {
            throw new Error("dev: unsafe values gamma");
        }
        if (D <= 10n ** 17n - 1n || D >= 10n ** 15n * 10n ** 18n + 1n) {
            throw new Error("dev: unsafe values D");
        }
        let lim_mul = 100n * 10n ** 18n;
        if (gamma > this.MAX_GAMMA_SMALL) {
            lim_mul = (lim_mul * this.MAX_GAMMA_SMALL) / gamma;
        }

        const y = this._newtonY(ann, gamma, x, D, i, lim_mul);
        const frac = (y * 10n ** 18n / D);
        if (frac < ((10n ** 36n / this.N_COINS) / lim_mul) || frac > (lim_mul / this.N_COINS)) {
            throw new Error("dev: unsafe values y");
        }

        return y;
    }

    public static getY(_ann: bigint, _gamma: bigint, _x: bigint[], _D: bigint, i: number): bigint[] {
        if (_ann <= this.MIN_A - 1n || _ann >= this.MAX_A + 1n) throw new Error("dev: unsafe values ann");
        if (_gamma <= this.MIN_GAMMA - 1n || _gamma >= this.MAX_GAMMA + 1n) throw new Error("dev: unsafe values gamma");
        if (_D <= 10n ** 17n - 1n || _D >= 10n ** 33n + 1n) throw new Error("dev: unsafe values D");
    
        let limMul = 100n * 10n ** 18n;
        if (_gamma > this.MAX_GAMMA_SMALL) {
            limMul = (limMul * this.MAX_GAMMA_SMALL) / _gamma;
        }
    
        const limMulSigned = this.convert(limMul, "int256");
        const ann = this.convert(_ann, "int256");
        const gamma = this.convert(_gamma, "int256");
        const D = this.convert(_D, "int256");
        const xj = this.convert(_x[1 - Number(i)], "int256");
        const gamma2 = gamma * gamma;
        const nCoins = this.convert(this.N_COINS, "int256");
    
        const y = D ** 2n / (xj * nCoins ** 2n);
    
        const k0_i = ((10n ** 18n * nCoins) * xj / D);
        if (k0_i < (10n ** 36n / limMulSigned) || k0_i > limMulSigned) {
            throw new Error("dev: unsafe values x[i]");
        }
    
        const annGamma2 = ann * gamma2;
        let a = 10n ** 32n;
    
        let b = (D * annGamma2 / 400000000n / xj)
            - MathUtils.convert(3n * 10n ** 32n, "int256")
            - (2n * gamma * 10n ** 14n);
    
        let c = (MathUtils.convert(3n, "int256") * 10n ** 32n)
            + (4n * gamma * 10n ** 14n)
            + (gamma2 / 10n ** 4n)
            + ((4n * annGamma2 / 400000000n) * xj / D)
            - (4n * annGamma2 / 400000000n);
    
        let d = -((10n ** 18n + gamma) ** 2n) / 10n ** 4n;
    
        let delta0 = (3n * a * c / b) - b;
        let delta1 = 3n * delta0 + b - ((27n * a ** 2n / b) * d / b);
    
        let divider = 1n;
        const threshold = this.min(this.min(this.abs(delta0), this.abs(delta1)), a);
        if (threshold > 10n ** 48n) divider = 10n ** 30n;
        else if (threshold > 10n ** 46n) divider = 10n ** 28n;
        else if (threshold > 10n ** 44n) divider = 10n ** 26n;
        else if (threshold > 10n ** 42n) divider = 10n ** 24n;
        else if (threshold > 10n ** 40n) divider = 10n ** 22n;
        else if (threshold > 10n ** 38n) divider = 10n ** 20n;
        else if (threshold > 10n ** 36n) divider = 10n ** 18n;
        else if (threshold > 10n ** 34n) divider = 10n ** 16n;
        else if (threshold > 10n ** 32n) divider = 10n ** 14n;
        else if (threshold > 10n ** 30n) divider = 10n ** 12n;
        else if (threshold > 10n ** 28n) divider = 10n ** 10n;
        else if (threshold > 10n ** 26n) divider = 10n ** 8n;
        else if (threshold > 10n ** 24n) divider = 10n ** 6n;
        else if (threshold > 10n ** 20n) divider = 10n ** 2n;
    
        a = a / divider;
        b = b / divider;
        c = c / divider;
        d = d / divider;
    
        delta0 = (3n * a * c / b) - b;
        delta1 = 3n * delta0 + b - ((27n * a ** 2n / b) * d / b);
    
        const sqrtArg = delta1 ** 2n + ((4n * delta0 ** 2n / b) * delta0);
    
        let sqrtVal = 0n;
        if (sqrtArg > 0n) {
            sqrtVal = this.convert(this.sqrt(this.convert(sqrtArg, "uint256")), "int256");
        } else {
            return [this._newtonY(_ann, _gamma, _x, _D, i, limMul), 0n];
        }
    
        let bCbrt = b > 0n
            ? this.convert(this.cbrt(this.convert(b, "uint256")), "int256")
            : -this.convert(this.cbrt(this.convert(-b, "uint256")), "int256");
    
        let secondCbrt = delta1 > 0n
            ? this.convert(this.cbrt(this.convert((delta1 + sqrtVal), "uint256") / 2n), "int256")
            : -this.convert(this.cbrt(this.convert((sqrtVal - delta1), "uint256") / 2n), "int256");
    
        const C1 = ((bCbrt ** 2n / 10n ** 18n) * secondCbrt) / 10n ** 18n;
        const root = (
            ((10n ** 18n * C1) - (10n ** 18n * b) - ((10n ** 18n * b) / C1) * delta0) / (3n * a)
        );


        const yOut = [
            this.convert(((((D ** 2n / xj) * root) / 4n) / 10n ** 18n), "uint256"),
            this.convert(root, "uint256")
        ];
        const frac = (yOut[0] * 10n ** 18n) / _D;
        if (frac < (10n ** 36n / this.N_COINS) / limMul || frac > limMul / this.N_COINS) {
            throw new Error("unsafe value for y");
        }
    
        return yOut;
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

    static newtonD(ann: bigint, gamma: bigint, x_unsorted: bigint[], k0_prev = 0n): bigint {
        if (ann <= this.MIN_A - 1n && ann >= this.MAX_A + 1n) {
            throw new Error("dev: unsafe values ann");
        }
        if (gamma <= this.MIN_GAMMA - 1n && gamma >= this.MAX_GAMMA + 1n) {
            throw new Error("dev: unsafe values gamma");
        }
        
        let x: bigint[] = x_unsorted;
        if (x[0] < x[1]) {
            x = [x_unsorted[1], x_unsorted[0]];
        }

        if (x[0] <= 10n ** 9n - 1n && x[0] >= 10n ** 15n * 10n ** 18n + 1n) {
            throw new Error("dev: unsafe values x[0]");
        }
        if (((x[1] * 10n ** 18n) / x[0]) <= 10n ** 14n - 1n) {
            throw new Error("dev: unsafe values x[1] (input)");
        }

        let S = x[0] + x[1];

        let D = 0n;

        if (k0_prev === 0n) {
            D = this.N_COINS * this.sqrt(x[0] * x[1]);
        } else {
            D = this.sqrt(((4n * x[1]) * x[1] / k0_prev) * 10n ** 18n);
            if (S < D) {
                D = S;
            }
        }

        let __g1k0 = gamma + 10n ** 18n;
        let diff = 0n;

        for (let i = 0; i < 255; i++) {
            let D_prev = D;
            if (D <= 0n) {
                throw new Error("dev: D == 0");
            }
            let k0 = (((10n ** 18n * this.N_COINS ** 2n * x[0]) / D) * x[1]) / D;

            let _g1k0 = __g1k0;
            if (_g1k0 > k0) {
                _g1k0 = _g1k0 - k0 + 1n;
            } else {
                _g1k0 = k0 - _g1k0 + 1n;
            }

            const mul1 = (((((10n ** 18n * D) / gamma) * _g1k0) / gamma) * _g1k0 * this.A_MULTIPLIER) / ann;
            const mul2 = (2n * 10n ** 18n * this.N_COINS * k0) / _g1k0;

            const neg_fprime = S + (S * mul2 / 10n ** 18n) + (mul1 * this.N_COINS / k0) - (mul2 * D) / 10n ** 18n;

            let D_plus = (D * (neg_fprime + S) / neg_fprime);
            let D_minus = D * D / neg_fprime;

            if (k0 < 10n ** 18n) {
                D_minus += ((D * (mul1 / neg_fprime) / 10n ** 18n) * (10n ** 18n - k0)) / k0;
            } else {
                D_minus -= ((D * (mul1 / neg_fprime) / 10n ** 18n) * (k0 - 10n ** 18n)) / k0;
            }

            if (D_plus > D_minus) {
                D = D_plus - D_minus;
            } else {
                D = (D_minus - D_plus) /  2n;
            }

            if (D > D_prev) {
                diff = D - D_prev;
            } else {
                diff = D_prev - D;
            }

            if (diff * 10n ** 14n < this.max(10n ** 16n, D)) {
                x.forEach((_x: bigint) => {
                    let frac = (_x * 10n ** 18n / D);
                    if (frac <= (10n ** 16n / this.N_COINS - 1n) && frac >= (10n ** 20n / this.N_COINS + 1n)) {
                        throw new Error("unsafe value x[i]");
                    }
                })
                return D;
            }
        }
        throw new Error("dev: did not converge");
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
}

export class TwoCryptoSwap {
    private static readonly PRECISION = 10n ** 18n;
    private static readonly FEE_DENOMINATOR = 10n ** 10n;

    static xp: bigint[];
    static precisions: bigint[];
    static price_scale: bigint;
    static A: bigint;
    static D: bigint;
    static GAMMA: bigint;
    static FUTURE_A_GAMMA_TIME: bigint;
    static MID_FEE: bigint;
    static OUT_FEE: bigint;
    static FEE_GAMMA: bigint;

    constructor(xp_: bigint[], precisions_: bigint[], price_scale_: bigint, a: bigint, d: bigint, gamma: bigint, future_a_gamma_time: bigint, mid_fee: bigint, out_fee: bigint, fee_gamma: bigint) {
        TwoCryptoSwap.xp = xp_;
        TwoCryptoSwap.precisions = precisions_;
        TwoCryptoSwap.price_scale = price_scale_;
        TwoCryptoSwap.A = a;
        TwoCryptoSwap.D = d;
        TwoCryptoSwap.GAMMA = gamma;
        TwoCryptoSwap.FUTURE_A_GAMMA_TIME = future_a_gamma_time;
        TwoCryptoSwap.MID_FEE = mid_fee;
        TwoCryptoSwap.OUT_FEE = out_fee;
        TwoCryptoSwap.FEE_GAMMA = fee_gamma;
    }

    static calculateDRamp(A: bigint, gamma: bigint, xp: bigint[], precisions: bigint[], price_scale: bigint): bigint {
        const timestamp = BigInt(Math.floor(Date.now() / 1000)); // Current time in seconds as a BigInt
        let D = TwoCryptoSwap.D;
        if (TwoCryptoSwap.FUTURE_A_GAMMA_TIME > timestamp) {
            let _xp = xp;
            _xp[0] = _xp[0] * precisions[0];
            _xp[1] = (_xp[1] * price_scale * precisions[1]) / TwoCryptoSwap.PRECISION
            D = MathUtils.newtonD(A, gamma, _xp);
        }

        return D;
    }

    static getDyNoFee(i: number, j: number, dx: bigint): [bigint, bigint[]] {
        if (i === j && i >= Number(MathUtils.N_COINS) && j >= Number(MathUtils.N_COINS)) {
            throw new Error("dev: coin index out of range");
        }
        if (dx <= 0n) {
            throw new Error("dev: do not exchange 0 coins");
        }
        let D = this.calculateDRamp(TwoCryptoSwap.A, TwoCryptoSwap.D, TwoCryptoSwap.xp, TwoCryptoSwap.precisions, TwoCryptoSwap.price_scale);
        let _xp = TwoCryptoSwap.xp;
        _xp[i] += dx;
        _xp = [
            _xp[0] * TwoCryptoSwap.precisions[0],
            (_xp[1] * TwoCryptoSwap.price_scale * TwoCryptoSwap.precisions[1]) / TwoCryptoSwap.PRECISION,
        ]
        const y_out = MathUtils.getY(TwoCryptoSwap.A, TwoCryptoSwap.GAMMA, _xp, D, j);
        let dy = _xp[j] - y_out[0] - 1n; 
        _xp[j] = y_out[0];
        if (j > 0) {
            dy = (dy * TwoCryptoSwap.PRECISION) / TwoCryptoSwap.price_scale;
        }
        dy /= TwoCryptoSwap.precisions[j];
        return [dy, _xp];
    }

    static feeCalculate(xp: bigint[]): bigint {
        let f = xp[0] + xp[1];
        f = this.FEE_GAMMA * (10n ** 18n) / (this.FEE_GAMMA + (10n ** 18n) - (((((10n ** 18n) * (MathUtils.N_COINS ** MathUtils.N_COINS) * xp[0]) / f) * xp[1]) / f));
        return (this.MID_FEE * f + this.OUT_FEE * (10n ** 18n - f)) / (10n ** 18n);
    }

    public getDy(i: number, j: number, dx: bigint): bigint {
        let [dy, _xp] = TwoCryptoSwap.getDyNoFee(i, j, dx);
        const fee = TwoCryptoSwap.feeCalculate(_xp);
        dy -= (dy * fee / TwoCryptoSwap.FEE_DENOMINATOR);
        return dy;
    }
}

async function main() {
    const provider = new ethers.providers.JsonRpcProvider("https://ethereum-rpc.publicnode.com");
    const poolAddress = "0xc907ba505c2e1cbc4658c395d4a2c7e6d2c32656";
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

    async function fetchGetDy(i: number, j: number, amountIn: bigint) {
        const dy = await poolContract.get_dy(i, j, amountIn);
        return BigInt(dy.toString());
    }


    let price_scale = await fetchPriceScale();
    const A = await fetchA();
    const D = await fetchD();
    const gamma = await fetchGamma();
    const future_a_gamma_time = await fetchFutureAGammaTime();
    const mid_fee = await fetchMidFee();
    const out_fee = await fetchOutFee();
    const fee_gamma = await fetchFeeGamma();

    let xp: bigint[] = [];
    xp[0] = await fetchBalances(0n);
    xp[1] = await fetchBalances(1n);

    const twoCryptoSwap = new TwoCryptoSwap(xp, [1n, 1n], price_scale, A, D, gamma, future_a_gamma_time, mid_fee, out_fee, fee_gamma);

    const i = 0; // Token in: WETH
    const j = 1; // Token out: CVX
    const amountIn = 1000000000000000000n; 

    try {
        const dy = twoCryptoSwap.getDy(i, j, amountIn);
        console.log(`Amount out (dy) calculated: ${dy.toString()}`);
        console.log("Amount out (dy) from contract: ", (await fetchGetDy(i, j, amountIn)).toString());
    } catch (error) {
        console.error("Error calculating dy:", error);
    }
}
main();