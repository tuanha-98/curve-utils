const { ethers } = require("ethers");
import JSBI from "jsbi";

export class MathUtils {
    static readonly N_COINS = 3n;
    private static readonly A_MULTIPLIER = 10000n;
    private static readonly MIN_GAMMA = 10n ** 10n;
    private static readonly MAX_GAMMA = 2n * 10n ** 16n;
    private static readonly MIN_A = (this.N_COINS ** this.N_COINS * this.A_MULTIPLIER / 100n) ;
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

    private static _newtonY(ann: bigint, gamma: bigint, x: bigint[], D: bigint, i: number): bigint {
        let frac = 0n;
        for (let k = 0; k < 3; k++) {
            if (k != i) {
                frac = (x[k] * 10n ** 18n) / D;
                if (frac <= 10n ** 16n - 1n || frac >= 10n ** 20n * this.N_COINS + 1n) {
                    throw new Error("dev: unsafe values x[i]");
                }
            }
        }

        let y = D / this.N_COINS;
        let k0_i = 10n ** 18n;
        let S_i = 0n;
        let x_sorted = x;
        x_sorted[i] = 0n;
        x_sorted = x_sorted.sort((a: bigint, b: bigint) => a < b ? 1 : a > b ? -1 : 0);
        
        const convergence_limit = this.max(this.max((x_sorted[0] / 10n ** 14n), (D / 10n ** 14n)), 100n);

        for (let j = 2; j <= Number(this.N_COINS); j++) {
            const _x = x_sorted[Number(this.N_COINS) - j];
            y = (y * D) / (_x * this.N_COINS);  // Small _x first
            S_i += _x;
        }

        for (let  j = 0; j < Number(this.N_COINS); j++) {
            k0_i = k0_i * x_sorted[j] * this.N_COINS / D
        }

        let diff = 0n;
        let y_prev = 0n;
        let k0 = 0n;
        let S = 0n;
        let g1k0 = 0n;
        let mul1 = 0n;
        let mul2 = 0n;
        let yfprime = 0n;
        let dyfprime = 0n;
        let fprime = 0n;
        let y_minus = 0n;
        let y_plus = 0n;

        for (let j = 0; j < 255; ++j) {
            y_prev = y;
            
            k0 = k0_i * y * this.N_COINS / D;
            S = S_i + y;

            g1k0 = gamma + 10n ** 18n;
            if (g1k0 > k0) {
                g1k0 = g1k0 - k0 + 1n;
            } else {
                g1k0 = k0 - g1k0 + 1n;
            }

            mul1 = (((((10n ** 18n * D) / gamma) * g1k0) / gamma) * g1k0 * this.A_MULTIPLIER) / ann;
            mul2 = 10n ** 18n + (2n * 10n ** 18n) * k0 / g1k0;

            yfprime = 10n ** 18n * y + S * mul2 + mul1;
            dyfprime = D * mul2;

            if (yfprime < dyfprime) {
                y = y_prev / 2n;
                continue;
            } else {
                yfprime -= dyfprime;
            }
            fprime = yfprime / y;
            y_minus = mul1 / fprime;
            y_plus = (yfprime + 10n ** 18n * D) / fprime + y_minus * 10n ** 18n / k0;
            y_minus += 10n ** 18n * S / fprime;

            if (y_plus < y_minus) {
                y = y_prev / 2n;
            } else {
                y = y_plus - y_minus;
            }

            if (y > y_prev) {
                diff = y - y_prev;
            } else {
                diff = y_prev - y;
            }
            if (diff < this.max(convergence_limit, y / 10n ** 14n)) {
                frac = y * 10n ** 18n / D;
                if (frac <= 10n ** 16n - 1n || frac >= 10n ** 20n + 1n) {
                    throw new Error("dev: unsafe values y");
                }
                return y;
            }
        }
        throw new Error("dev: did not converge");
    }

    // public static newtonY(ann: bigint, gamma: bigint, x: bigint[], D: bigint, i: number): bigint {
    //     if (ann <= this.MIN_A - 1n || ann >= this.MAX_A + 1n) {
    //         throw new Error("dev: unsafe values ann");
    //     }
    //     if (gamma <= this.MIN_GAMMA - 1n || gamma >= this.MAX_GAMMA + 1n) {
    //         throw new Error("dev: unsafe values gamma");
    //     }
    //     if (D <= 10n ** 17n - 1n || D >= 10n ** 15n * 10n ** 18n + 1n) {
    //         throw new Error("dev: unsafe values D");
    //     }
    //     let lim_mul = 100n * 10n ** 18n;
    //     if (gamma > this.MAX_GAMMA_SMALL) {
    //         lim_mul = (lim_mul * this.MAX_GAMMA_SMALL) / gamma;
    //     }

    //     const y = this._newtonY(ann, gamma, x, D, i, lim_mul);
    //     const frac = (y * 10n ** 18n / D);
    //     if (frac < ((10n ** 36n / this.N_COINS) / lim_mul) || frac > (lim_mul / this.N_COINS)) {
    //         throw new Error("dev: unsafe values y");
    //     }

    //     return y;
    // }

    public static getY(_ann: bigint, _gamma: bigint, x: bigint[], _D: bigint, i: number): bigint[] {
        if (_ann <= this.MIN_A - 1n || _ann >= this.MAX_A + 1n) throw new Error("dev: unsafe values ann");
        if (_gamma <= this.MIN_GAMMA - 1n || _gamma >= this.MAX_GAMMA + 1n) throw new Error("dev: unsafe values gamma");
        if (_D <= 10n ** 17n - 1n || _D >= 10n ** 33n + 1n) throw new Error("dev: unsafe values D");
    
        let frac = 0n;
        for (let k = 0; k < 3; k++) {
            if (k != i) {
                frac = (x[k] * 10n ** 18n) / _D;
                if (frac <= 10n ** 16n - 1n || frac >= 10n ** 20n + 1n) {
                    throw new Error("dev: unsafe values x[i]");
                }
            }
        }

        let j = 0;
        let k = 0;

        if (i == 0) {
            j = 1;
            k = 2;
        } else if (i == 1) {
            j = 0;
            k = 2;
        } else if (i == 2) {
            j = 0;
            k = 1;
        }

        const ann = this.convert(_ann, "int256");
        const gamma = this.convert(_gamma, "int256");
        const D = this.convert(_D, "int256");
        const xj = this.convert(x[j], "int256");
        const xk = this.convert(x[k], "int256");
        const gamma2 = gamma * gamma;
    
        let a = 10n ** 36n / 27n;
    
        let b = (
            ((10n ** 36n / 9n) + (2n * 10n ** 18n * gamma) / 27n)
            - (((D * D / xj) * gamma2 * ann) / (27n ** 2n) / MathUtils.convert(this.A_MULTIPLIER, "int256") / xk)
        );
    
        let c = (
            ((10n ** 36n / 9n) + ((gamma * (gamma + 4n * 10n ** 18n))) / 27n)
            + ((((gamma2 * ((xj + xk) - D)) / D) * ann) / 27n) / MathUtils.convert(this.A_MULTIPLIER, "int256")
        );
    
        let d = (10n ** 18n + gamma) ** 2n / 27n;
    
        let d0 = MathUtils.abs((3n * a) * c / b - b);
    
        let divider = 0n;
        if (d0 > 10n ** 48n) divider = 10n ** 30n;
        else if (d0 > 10n ** 46n) divider = 10n ** 28n;
        else if (d0 > 10n ** 44n) divider = 10n ** 26n;
        else if (d0 > 10n ** 42n) divider = 10n ** 24n;
        else if (d0 > 10n ** 40n) divider = 10n ** 22n;
        else if (d0 > 10n ** 38n) divider = 10n ** 20n;
        else if (d0 > 10n ** 36n) divider = 10n ** 18n;
        else if (d0 > 10n ** 34n) divider = 10n ** 16n;
        else if (d0 > 10n ** 32n) divider = 10n ** 14n;
        else if (d0 > 10n ** 30n) divider = 10n ** 12n;
        else if (d0 > 10n ** 28n) divider = 10n ** 10n;
        else if (d0 > 10n ** 26n) divider = 10n ** 8n;
        else if (d0 > 10n ** 24n) divider = 10n ** 6n;
        else if (d0 > 10n ** 20n) divider = 10n ** 2n;
        else divider = 1n;
        
        let additional_prec = 0n;
        if (MathUtils.abs(a) > MathUtils.abs(b)) {
            additional_prec = MathUtils.abs(a / b);
            a = (a * additional_prec) / divider;
            b = (b * additional_prec) / divider;
            c = (c * additional_prec) / divider;
            d = (d * additional_prec) / divider;
        } else {
            additional_prec = MathUtils.abs(b / a);
            a = (a / additional_prec) / divider;
            b = (b / additional_prec) / divider;
            c = (c / additional_prec) / divider;
            d = (d / additional_prec) / divider;
        }

        let _3ac = 3n * a * c;
    
        let delta0 = (_3ac / b) - b;
        let delta1 = (3n * _3ac / b) - (2n * b) - (((27n * a ** 2n) / b) * d) / b;
    
        const sqrtArg = delta1 ** 2n + ((4n * delta0 ** 2n / b) * delta0);
    
        let sqrtVal = 0n;
        if (sqrtArg > 0n) {
            sqrtVal = this.convert(this.sqrt(this.convert(sqrtArg, "uint256")), "int256");
        } else {
            return [this._newtonY(_ann, _gamma, x, _D, i), 0n];
        }
    
        let bCbrt = b > 0n
            ? this.convert(this.cbrt(this.convert(b, "uint256")), "int256")
            : -this.convert(this.cbrt(this.convert(-b, "uint256")), "int256");
    
        let secondCbrt = delta1 > 0n
            ? this.convert(this.cbrt(this.convert((delta1 + sqrtVal), "uint256") / 2n), "int256")
            : -this.convert(this.cbrt(this.convert((sqrtVal - delta1), "uint256") / 2n), "int256");
    
        const C1 = ((bCbrt ** 2n / 10n ** 18n) * secondCbrt) / 10n ** 18n;

        const root_k0 = (b + b * delta0 / C1 - C1) / 3n
        const root = (((((D * D / 27n) / xk) * D) / xj) * root_k0) / a;

        const out = [
            this.convert(root, "uint256"),
            this.convert((10n ** 18n * root_k0) / a, "uint256")
        ];

        frac = (out[0] * 10n ** 18n) / _D;
        if (frac < (10n ** 16n - 1n) || frac >= 10n ** 20n + 1n) {
            throw new Error("unsafe value for y");
        }
    
        return out;
    }
    
    static geometricMean(_x: bigint[]): bigint {
        let prod = (((_x[0] * _x[1]) / 10n ** 18n) * _x[2]) / 10n ** 18n;
        if (prod == 0n) {
            return 0n;
        }
        return this.cbrt(prod);
    }

    static newtonD(ann: bigint, gamma: bigint, x_unsorted: bigint[], k0_prev = 0n): bigint {
        let x = x_unsorted.sort((a: bigint, b: bigint) => a < b ? 1 : a > b ? -1 : 0);

        if (x[0] >= MathUtils.convert(-1n, "uint256") / 10n ** 18n * this.N_COINS ** this.N_COINS) {
            throw new Error("dev: out of limits");
        }

        if (x[0] <= 0n) {
            throw new Error("dev: empty pool");
        }

        let S = x[0] + x[1] + x[2];
        let D = 0n;

        if (k0_prev === 0n) {
            D = this.N_COINS * this.geometricMean(x);
        } else {
            if (S > 10n ** 36n) {
                D = this.cbrt(((((x[0] * x[1]) / 10n ** 36n) * x[2]) / k0_prev) * 27n * 10n ** 12n);
            } else if (S > 10n ** 24n) {
                D = this.cbrt(((((x[0] * x[1]) / 10n ** 24n) * x[2]) / k0_prev) * 27n * 10n ** 6n);
            } else {
                D = this.cbrt(((((x[0] * x[1]) / 10n ** 18n) * x[2]) / k0_prev) * 27n);
            }
        }

        let k0 = 0n;
        let _g1k0 = 0n;
        let mul1 = 0n;
        let mul2 = 0n;
        let neg_fprime = 0n;
        let D_plus = 0n;
        let D_minus = 0n;
        let D_prev = 0n;

        let diff = 0n;
        let frac = 0n;

        for (let i = 0; i < 255; i++) {
            D_prev = D;
            
            k0 = ((((((((10n ** 18n * x[0]) * this.N_COINS) / D) * x[1]) * this.N_COINS) / D) * x[2]) * this.N_COINS) / D;

            _g1k0 = gamma + 10n ** 18n;
            if (_g1k0 > k0) {
                _g1k0 = _g1k0 - k0 + 1n;
            } else {
                _g1k0 = k0 - _g1k0 + 1n;
            }

            mul1 = (((((10n ** 18n * D) / gamma) * _g1k0) / gamma) * _g1k0 * this.A_MULTIPLIER) / ann;
            mul2 = (2n * 10n ** 18n * this.N_COINS * k0) / _g1k0;

            neg_fprime = (S + (S * mul2 / 10n ** 18n)) + (mul1 * this.N_COINS / k0) - (mul2 * D) / 10n ** 18n;

            D_plus = (D * (neg_fprime + S) / neg_fprime);
            D_minus = D * D / neg_fprime;

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
                    frac = (_x * 10n ** 18n / D);
                    if (frac < (10n ** 16n - 1n) && frac >= (10n ** 20n + 1n)) {
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

    static reductionCoefficient(x: bigint[], fee_gamma: bigint): bigint {
        let S = x[0] + x[1] + x[2];
        let K = 10n ** 18n * this.N_COINS * x[0] / S;
        K = (K * this.N_COINS * x[1]) / S;
        K = (K * this.N_COINS * x[2]) / S;

        if (fee_gamma > 0n) {
            K = fee_gamma * 10n ** 18n / (fee_gamma + 10n ** 18n - K);
        }
        return K;
    }
}

export class TricryptoSwap {
    private static readonly PRECISION = 10n ** 18n;
    private static readonly FEE_DENOMINATOR = 10n ** 10n;

    static xp: bigint[];
    static precisions: bigint[];
    static price_scale: bigint[];
    static A: bigint;
    static D: bigint;
    static GAMMA: bigint;
    static FUTURE_A_GAMMA_TIME: bigint;
    static MID_FEE: bigint;
    static OUT_FEE: bigint;
    static FEE_GAMMA: bigint;

    constructor(xp_: bigint[], precisions_: bigint[], price_scale_: bigint[], a: bigint, d: bigint, gamma: bigint, future_a_gamma_time: bigint, mid_fee: bigint, out_fee: bigint, fee_gamma: bigint) {
        TricryptoSwap.xp = xp_;
        TricryptoSwap.precisions = precisions_;
        TricryptoSwap.price_scale = price_scale_;
        TricryptoSwap.A = a;
        TricryptoSwap.D = d;
        TricryptoSwap.GAMMA = gamma;
        TricryptoSwap.FUTURE_A_GAMMA_TIME = future_a_gamma_time;
        TricryptoSwap.MID_FEE = mid_fee;
        TricryptoSwap.OUT_FEE = out_fee;
        TricryptoSwap.FEE_GAMMA = fee_gamma;
    }

    static calculateDRamp(A: bigint, gamma: bigint, xp: bigint[], precisions: bigint[], price_scale: bigint[]): bigint {
        const timestamp = BigInt(Math.floor(Date.now() / 1000)); // Current time in seconds as a BigInt
        let D = TricryptoSwap.D;
        if (TricryptoSwap.FUTURE_A_GAMMA_TIME > timestamp) {
            let _xp = xp;
            _xp[0] = _xp[0] * precisions[0];
            for (let k = 0; k < Number(MathUtils.N_COINS) - 1; k++) {
                _xp[k + 1] = (_xp[k + 1] * price_scale[k] * precisions[k + 1]) / TricryptoSwap.PRECISION;
            }
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
        let D = this.calculateDRamp(TricryptoSwap.A, TricryptoSwap.D, TricryptoSwap.xp, TricryptoSwap.precisions, TricryptoSwap.price_scale);
        let _xp = TricryptoSwap.xp;
        _xp[i] += dx;
        _xp[0] *= TricryptoSwap.precisions[0];

        for (let k = 0; k < Number(MathUtils.N_COINS) - 1; k++) {
            _xp[k + 1] = _xp[k + 1] * TricryptoSwap.price_scale[k] * TricryptoSwap.precisions[k + 1] / TricryptoSwap.PRECISION;
        }

        const y_out = MathUtils.getY(TricryptoSwap.A, TricryptoSwap.GAMMA, _xp, D, j);
        let dy = _xp[j] - y_out[0] - 1n; 
        _xp[j] = y_out[0];
        if (j > 0) {
            dy = (dy * TricryptoSwap.PRECISION) / TricryptoSwap.price_scale[j - 1];
        }
        dy /= TricryptoSwap.precisions[j];
        return [dy, _xp];
    }

    static feeCalculate(xp: bigint[]): bigint {
        let f = MathUtils.reductionCoefficient(xp, TricryptoSwap.FEE_GAMMA);
        return (TricryptoSwap.MID_FEE * f + TricryptoSwap.OUT_FEE * (10n ** 18n - f)) / 10n ** 18n;
    }

    public getDy(i: number, j: number, dx: bigint): bigint {
        let [dy, _xp] = TricryptoSwap.getDyNoFee(i, j, dx);
        const fee = TricryptoSwap.feeCalculate(_xp);
        dy -= (dy * fee / TricryptoSwap.FEE_DENOMINATOR);
        return dy;
    }
}

async function main() {
    const provider = new ethers.providers.JsonRpcProvider("https://ethereum-rpc.publicnode.com");
    const poolAddress = "0x7F86Bf177Dd4F3494b841a37e810A34dD56c829B";
    const poolAbi = [
        "function balances(uint256) view returns (uint256)",
        "function price_scale(uint256) view returns (uint256)",
        "function gamma() view returns (uint256)",
        "function fee() view returns (uint256)",
        "function D() view returns (uint256)",
        "function A() view returns (uint256)",
        "function future_A_gamma_time() view returns (uint256)",
        "function mid_fee() view returns (uint256)",
        "function out_fee() view returns (uint256)",
        "function fee_gamma() view returns (uint256)",
        "function get_dy(uint256 i, uint256 j, uint256 dx) view returns (uint256)",
        "function precisions() view returns (uint256[3])",
    ];

    const poolContract = new ethers.Contract(poolAddress, poolAbi, provider);

    async function fetchBalances(index: bigint) {
        const balance = await poolContract.balances(index);
        return BigInt(balance.toString());
    }

    async function fetchPriceScale(index: bigint) {
        const priceScale = await poolContract.price_scale(index);
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

    async function fetchPrecisions() {
        const precisions = await poolContract.precisions();
        return precisions.map((p: any) => BigInt(p.toString()));
    }


    const A = await fetchA();
    const D = await fetchD();
    const gamma = await fetchGamma();
    const future_a_gamma_time = await fetchFutureAGammaTime();
    const mid_fee = await fetchMidFee();
    const out_fee = await fetchOutFee();
    const fee_gamma = await fetchFeeGamma();
    
    let price_scale: bigint[] = []
    price_scale[0] = await fetchPriceScale(0n);
    price_scale[1] = await fetchPriceScale(1n);

    let xp: bigint[] = [];
    xp[0] = await fetchBalances(0n);
    xp[1] = await fetchBalances(1n);
    xp[2] = await fetchBalances(2n);

    let precisions: bigint[] = await fetchPrecisions();

    const tricryptoSwap = new TricryptoSwap(xp, precisions, price_scale, A, D, gamma, future_a_gamma_time, mid_fee, out_fee, fee_gamma);

    const i = 0; // Token in: WETH
    const j = 1; // Token out: CVX
    const amountIn = 1000000000n; 

    try {
        const dy = tricryptoSwap.getDy(i, j, amountIn);
        console.log(`Amount out (dy) calculated: ${dy.toString()}`);
        console.log("Amount out (dy) from contract: ", (await fetchGetDy(i, j, amountIn)).toString());
    } catch (error) {
        console.error("Error calculating dy:", error);
    }
}
main();