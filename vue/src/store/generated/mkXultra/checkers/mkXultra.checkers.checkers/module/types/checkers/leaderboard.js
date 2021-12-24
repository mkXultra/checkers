/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "mkXultra.checkers.checkers";
const baseLeaderboard = { winners: "" };
export const Leaderboard = {
    encode(message, writer = Writer.create()) {
        if (message.winners !== "") {
            writer.uint32(10).string(message.winners);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseLeaderboard };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.winners = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseLeaderboard };
        if (object.winners !== undefined && object.winners !== null) {
            message.winners = String(object.winners);
        }
        else {
            message.winners = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.winners !== undefined && (obj.winners = message.winners);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseLeaderboard };
        if (object.winners !== undefined && object.winners !== null) {
            message.winners = object.winners;
        }
        else {
            message.winners = "";
        }
        return message;
    },
};
