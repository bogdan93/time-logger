import { eifActions } from "../../store";
import { getDateFrom } from "../pages/calendar/utils";
import MailForm from './MailForm.svelte';
import type { MailReportForm } from "./types";

export function openMailReport () {
  eifActions.pushLayoutBlock({
    type: 'modal',
    title: 'Mail report',
    dataPath: 'mail-report',
    component: MailForm,
    actions: [
      {
        title: 'cancel',
        async action(block) {
          eifActions.popLayoutBlock(block.blockId, 'deleteDataAtPath');
        }
      },
      {
        title: 'download',
        validateBeforeExecute: true,
        async action(block) {
          const mailForm = eifActions.getDataAtPath<MailReportForm>(block.dataPath);
          const { month, year } = eifActions.getData(state => state.selectedMonth);
          const data = getDateFrom(1, month, year);

          // @ts-ignore
          window.html2pdf(document.getElementById('to-pdf'), {
            filename: `Raport activitate(${mailForm?.invoiceName}) - ${data.format('MMM YYYY')}.pdf`,
            html2canvas: {
              scale: 2,
              letterRendering: true,
            },
            pageBreak: {
              mode: ['legacy'],
            },
            jsPDF: { unit: 'mm', format: 'a4', orientation: 'portrait' }
          });
        }
      },
    ],
  });
}
