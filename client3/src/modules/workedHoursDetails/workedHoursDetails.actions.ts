import api from "../../api";
import workedHoursDetailsSlice, { type WorkedHoursDetails } from "./workedHoursDetails.slice";

export async function loadWorkhourDetails() {
  try {
    const workhour_details = await api.records.getFullList('workhour_details', 200, {
      sort: '-created',
    }) as unknown as WorkedHoursDetails[] | null;

    if (!workhour_details) {
      return;
    }

    (workhour_details || []).forEach(workedHoursDetailsSlice.actions.set);

  } catch (error) {
    console.log(error);
    console.error("ERROR loading workhour_details");
  }
}
