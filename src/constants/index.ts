import { exp10 } from "../utils";

export const A_MULTIPLIER = 10000n;
export const PRECISION = exp10(18);
export const FEE_DENOMINATOR = exp10(10);

export const MIN_GAMMA = exp10(10);
export const MAX_GAMMA = 2n * exp10(16);