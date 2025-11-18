import api from "../api";
import { eifActions } from "../store";
import { goToLogin } from "./auth/auth.actions";
import loggedInUserSlice from "./auth/loggedInUser.slice";
import { loadCompanies } from "./companies/companies.actions";
import companiesSlice from "./companies/companies.slice";
import { loadCompaniesContractDetails } from "./contractDetails/contractDetails.actions";
import companiesContractDetailsSlice from "./contractDetails/contractDetails.slice";
import { openMailReport } from "./mailReport/mail.actions";
import { openOdooExport } from "./odoo/odoo.actions";
import Calendar from './pages/calendar/Calendar.svelte';
import selectedMonthSlice from "./pages/calendar/selectedMonth.slice";
import workedHoursSlice from "./pages/calendar/workedHours.slice";
import workedHoursCopySlice from "./pages/calendar/workedHoursCopy.slice";
import { loadProjects } from "./projects/projects.actions";
import projectsSlice from "./projects/projects.slice";
import { loadSignatures } from "./signatures/signatures.actions";
import signaturesSlice from "./signatures/signatures.slice";
import { loadWorkhourDetails } from "./workedHoursDetails/workedHoursDetails.actions";
import workedHoursDetailsSlice from "./workedHoursDetails/workedHoursDetails.slice";

export async function appInit() {
  loggedInUserSlice.actions.set({ email: api.authStore.model.email });

  await Promise.allSettled([
    loadProjects(),
    loadWorkhourDetails(),
    loadSignatures(),
    loadCompanies(),
    loadCompaniesContractDetails(),
  ]);

  eifActions.replaceLayoutBlock({
    type: "page",
    title: "Calendar",
    dataPath: "",
    component: Calendar,
    actions: [
      {
        title: 'Mail Report',
        action: openMailReport,
      },
      {
        title: 'Odoo Report',
        action: openOdooExport
      },
      {
        title: "Logout",
        action: goToLogin,
      },
    ],
  }, 'deleteDataAtPath');
}

export function appDestroy() {
  api.authStore.clear();
  loggedInUserSlice.actions.reset();
  projectsSlice.actions.reset();
  workedHoursDetailsSlice.actions.reset();
  selectedMonthSlice.actions.reset();
  workedHoursSlice.actions.reset();
  workedHoursCopySlice.actions.reset();
  signaturesSlice.actions.reset();
  companiesSlice.actions.reset();
  companiesContractDetailsSlice.actions.reset();
}
