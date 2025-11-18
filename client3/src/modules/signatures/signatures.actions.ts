import api from "../../api";
import signaturesSlice from "./signatures.slice";
import type { Signature } from "./types";

export async function loadSignatures() {
  try {
    const records = await api.records.getFullList('signatures', 200, {
      sort: '-created',
    }) as unknown as Signature[];

    (records || []).forEach(signaturesSlice.actions.set);

  } catch (error) {
    console.log(error);
    console.error("ERROR loading signatures");
  }
}
