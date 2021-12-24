import { Reader, Writer } from "protobufjs/minimal";
import { NextGame } from "../checkers/next_game";
import { StoredGame } from "../checkers/stored_game";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { PlayerInfo } from "../checkers/player_info";
export declare const protobufPackage = "mkXultra.checkers.checkers";
export interface QueryGetNextGameRequest {
}
export interface QueryGetNextGameResponse {
    NextGame: NextGame | undefined;
}
export interface QueryGetStoredGameRequest {
    index: string;
}
export interface QueryGetStoredGameResponse {
    storedGame: StoredGame | undefined;
}
export interface QueryAllStoredGameRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllStoredGameResponse {
    storedGame: StoredGame[];
    pagination: PageResponse | undefined;
}
export interface QueryCanPlayMoveRequest {
    idValue: string;
    player: string;
    fromX: number;
    fromY: number;
    toX: number;
    toY: number;
}
export interface QueryCanPlayMoveResponse {
    possible: boolean;
    reason: string;
}
export interface QueryGetPlayerInfoRequest {
    index: string;
}
export interface QueryGetPlayerInfoResponse {
    playerInfo: PlayerInfo | undefined;
}
export interface QueryAllPlayerInfoRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllPlayerInfoResponse {
    playerInfo: PlayerInfo[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetNextGameRequest: {
    encode(_: QueryGetNextGameRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetNextGameRequest;
    fromJSON(_: any): QueryGetNextGameRequest;
    toJSON(_: QueryGetNextGameRequest): unknown;
    fromPartial(_: DeepPartial<QueryGetNextGameRequest>): QueryGetNextGameRequest;
};
export declare const QueryGetNextGameResponse: {
    encode(message: QueryGetNextGameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetNextGameResponse;
    fromJSON(object: any): QueryGetNextGameResponse;
    toJSON(message: QueryGetNextGameResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetNextGameResponse>): QueryGetNextGameResponse;
};
export declare const QueryGetStoredGameRequest: {
    encode(message: QueryGetStoredGameRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetStoredGameRequest;
    fromJSON(object: any): QueryGetStoredGameRequest;
    toJSON(message: QueryGetStoredGameRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetStoredGameRequest>): QueryGetStoredGameRequest;
};
export declare const QueryGetStoredGameResponse: {
    encode(message: QueryGetStoredGameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetStoredGameResponse;
    fromJSON(object: any): QueryGetStoredGameResponse;
    toJSON(message: QueryGetStoredGameResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetStoredGameResponse>): QueryGetStoredGameResponse;
};
export declare const QueryAllStoredGameRequest: {
    encode(message: QueryAllStoredGameRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllStoredGameRequest;
    fromJSON(object: any): QueryAllStoredGameRequest;
    toJSON(message: QueryAllStoredGameRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllStoredGameRequest>): QueryAllStoredGameRequest;
};
export declare const QueryAllStoredGameResponse: {
    encode(message: QueryAllStoredGameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllStoredGameResponse;
    fromJSON(object: any): QueryAllStoredGameResponse;
    toJSON(message: QueryAllStoredGameResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllStoredGameResponse>): QueryAllStoredGameResponse;
};
export declare const QueryCanPlayMoveRequest: {
    encode(message: QueryCanPlayMoveRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryCanPlayMoveRequest;
    fromJSON(object: any): QueryCanPlayMoveRequest;
    toJSON(message: QueryCanPlayMoveRequest): unknown;
    fromPartial(object: DeepPartial<QueryCanPlayMoveRequest>): QueryCanPlayMoveRequest;
};
export declare const QueryCanPlayMoveResponse: {
    encode(message: QueryCanPlayMoveResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryCanPlayMoveResponse;
    fromJSON(object: any): QueryCanPlayMoveResponse;
    toJSON(message: QueryCanPlayMoveResponse): unknown;
    fromPartial(object: DeepPartial<QueryCanPlayMoveResponse>): QueryCanPlayMoveResponse;
};
export declare const QueryGetPlayerInfoRequest: {
    encode(message: QueryGetPlayerInfoRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPlayerInfoRequest;
    fromJSON(object: any): QueryGetPlayerInfoRequest;
    toJSON(message: QueryGetPlayerInfoRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetPlayerInfoRequest>): QueryGetPlayerInfoRequest;
};
export declare const QueryGetPlayerInfoResponse: {
    encode(message: QueryGetPlayerInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPlayerInfoResponse;
    fromJSON(object: any): QueryGetPlayerInfoResponse;
    toJSON(message: QueryGetPlayerInfoResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetPlayerInfoResponse>): QueryGetPlayerInfoResponse;
};
export declare const QueryAllPlayerInfoRequest: {
    encode(message: QueryAllPlayerInfoRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllPlayerInfoRequest;
    fromJSON(object: any): QueryAllPlayerInfoRequest;
    toJSON(message: QueryAllPlayerInfoRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllPlayerInfoRequest>): QueryAllPlayerInfoRequest;
};
export declare const QueryAllPlayerInfoResponse: {
    encode(message: QueryAllPlayerInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllPlayerInfoResponse;
    fromJSON(object: any): QueryAllPlayerInfoResponse;
    toJSON(message: QueryAllPlayerInfoResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllPlayerInfoResponse>): QueryAllPlayerInfoResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a nextGame by index. */
    NextGame(request: QueryGetNextGameRequest): Promise<QueryGetNextGameResponse>;
    /** Queries a storedGame by index. */
    StoredGame(request: QueryGetStoredGameRequest): Promise<QueryGetStoredGameResponse>;
    /** Queries a list of storedGame items. */
    StoredGameAll(request: QueryAllStoredGameRequest): Promise<QueryAllStoredGameResponse>;
    /** Queries a list of canPlayMove items. */
    CanPlayMove(request: QueryCanPlayMoveRequest): Promise<QueryCanPlayMoveResponse>;
    /** Queries a playerInfo by index. */
    PlayerInfo(request: QueryGetPlayerInfoRequest): Promise<QueryGetPlayerInfoResponse>;
    /** Queries a list of playerInfo items. */
    PlayerInfoAll(request: QueryAllPlayerInfoRequest): Promise<QueryAllPlayerInfoResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    NextGame(request: QueryGetNextGameRequest): Promise<QueryGetNextGameResponse>;
    StoredGame(request: QueryGetStoredGameRequest): Promise<QueryGetStoredGameResponse>;
    StoredGameAll(request: QueryAllStoredGameRequest): Promise<QueryAllStoredGameResponse>;
    CanPlayMove(request: QueryCanPlayMoveRequest): Promise<QueryCanPlayMoveResponse>;
    PlayerInfo(request: QueryGetPlayerInfoRequest): Promise<QueryGetPlayerInfoResponse>;
    PlayerInfoAll(request: QueryAllPlayerInfoRequest): Promise<QueryAllPlayerInfoResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
