/* eslint-disable */
import { NextGame } from "../checkers/next_game";
import { StoredGame } from "../checkers/stored_game";
import { PlayerInfo } from "../checkers/player_info";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "mkXultra.checkers.checkers";
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.nextGame !== undefined) {
            NextGame.encode(message.nextGame, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.storedGameList) {
            StoredGame.encode(v, writer.uint32(18).fork()).ldelim();
        }
        for (const v of message.playerInfoList) {
            PlayerInfo.encode(v, writer.uint32(26).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.storedGameList = [];
        message.playerInfoList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.nextGame = NextGame.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.storedGameList.push(StoredGame.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.playerInfoList.push(PlayerInfo.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.storedGameList = [];
        message.playerInfoList = [];
        if (object.nextGame !== undefined && object.nextGame !== null) {
            message.nextGame = NextGame.fromJSON(object.nextGame);
        }
        else {
            message.nextGame = undefined;
        }
        if (object.storedGameList !== undefined && object.storedGameList !== null) {
            for (const e of object.storedGameList) {
                message.storedGameList.push(StoredGame.fromJSON(e));
            }
        }
        if (object.playerInfoList !== undefined && object.playerInfoList !== null) {
            for (const e of object.playerInfoList) {
                message.playerInfoList.push(PlayerInfo.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.nextGame !== undefined &&
            (obj.nextGame = message.nextGame
                ? NextGame.toJSON(message.nextGame)
                : undefined);
        if (message.storedGameList) {
            obj.storedGameList = message.storedGameList.map((e) => e ? StoredGame.toJSON(e) : undefined);
        }
        else {
            obj.storedGameList = [];
        }
        if (message.playerInfoList) {
            obj.playerInfoList = message.playerInfoList.map((e) => e ? PlayerInfo.toJSON(e) : undefined);
        }
        else {
            obj.playerInfoList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.storedGameList = [];
        message.playerInfoList = [];
        if (object.nextGame !== undefined && object.nextGame !== null) {
            message.nextGame = NextGame.fromPartial(object.nextGame);
        }
        else {
            message.nextGame = undefined;
        }
        if (object.storedGameList !== undefined && object.storedGameList !== null) {
            for (const e of object.storedGameList) {
                message.storedGameList.push(StoredGame.fromPartial(e));
            }
        }
        if (object.playerInfoList !== undefined && object.playerInfoList !== null) {
            for (const e of object.playerInfoList) {
                message.playerInfoList.push(PlayerInfo.fromPartial(e));
            }
        }
        return message;
    },
};
