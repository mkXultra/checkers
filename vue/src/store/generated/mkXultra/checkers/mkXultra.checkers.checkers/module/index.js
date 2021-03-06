// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgPlayMove } from "./types/checkers/tx";
import { MsgCreatePost } from "./types/checkers/tx";
import { MsgRejectGame } from "./types/checkers/tx";
import { MsgCreateGame } from "./types/checkers/tx";
const types = [
    ["/mkXultra.checkers.checkers.MsgPlayMove", MsgPlayMove],
    ["/mkXultra.checkers.checkers.MsgCreatePost", MsgCreatePost],
    ["/mkXultra.checkers.checkers.MsgRejectGame", MsgRejectGame],
    ["/mkXultra.checkers.checkers.MsgCreateGame", MsgCreateGame],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgPlayMove: (data) => ({ typeUrl: "/mkXultra.checkers.checkers.MsgPlayMove", value: data }),
        msgCreatePost: (data) => ({ typeUrl: "/mkXultra.checkers.checkers.MsgCreatePost", value: data }),
        msgRejectGame: (data) => ({ typeUrl: "/mkXultra.checkers.checkers.MsgRejectGame", value: data }),
        msgCreateGame: (data) => ({ typeUrl: "/mkXultra.checkers.checkers.MsgCreateGame", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
