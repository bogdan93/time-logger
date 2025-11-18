import api from "../../api";
import companiesContractDetailsSlice, { type CompaniesContractDetails } from "./contractDetails.slice";

export async function loadCompaniesContractDetails() {
  try {
    const records = await api.records.getFullList('companies_contract_details', 200, {
      sort: '-created',
    }) as unknown as CompaniesContractDetails[];

    (records || []).forEach(companiesContractDetailsSlice.actions.set);
  } catch (error) {
    console.log(error);
    console.error("ERROR loading workhour_details");
  }
}
