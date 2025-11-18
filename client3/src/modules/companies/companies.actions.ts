import api from "../../api";
import companiesSlice from "./companies.slice";
import type { Company } from "./types";

export async function loadCompanies() {
  try {
    const records = await api.records.getFullList('companies', 200, {
      sort: '-created',
    }) as unknown as Company[];

    (records || []).forEach(companiesSlice.actions.set);

  } catch (error) {
    console.log(error);
    console.error("ERROR loading companies");
  }
}
