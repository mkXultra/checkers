import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "mkXultra.checkers.checkers";
export interface NextGame {
    creator: string;
    idValue: number;
    /** Will contain the index of the game at the head. */
    fifoHead: string;
    /** Will contain the index of the game at the tail. */
    fifoTail: string;
}
export declare const NextGame: {
    encode(message: NextGame, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): NextGame;
    fromJSON(object: any): NextGame;
    toJSON(message: NextGame): unknown;
    fromPartial(object: DeepPartial<NextGame>): NextGame;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
