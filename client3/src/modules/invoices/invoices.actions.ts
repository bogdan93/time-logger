import { eifActions } from "../../store";
import { getDateFrom } from "../pages/calendar/utils";
import InvoiceForm from './InvoiceForm.svelte';
import type { InvoiceFormData } from "./types";

export function openInvoiceForm() {
  const { month, year } = eifActions.getData(state => state.selectedMonth);
  const monthName = getDateFrom(1, month, year).format('MMMM');
  eifActions.pushLayoutBlock({
    type: 'modal',
    title: `Generate invoice for ${monthName}`,
    dataPath: 'invoice-generation-data',
    component: InvoiceForm,
    actions: [
      {
        title: 'cancel',
        action(block) {
          eifActions.popLayoutBlock(block.blockId, 'deleteDataAtPath');
        }
      },
      {
        title: 'preview',
        validateBeforeExecute: true,
        action() {
          eifActions.pushLayoutBlock({
            type: 'modal',
            title: `Preview invoice for ${monthName}`,
            dataPath: 'invoice-generation-data',
            component: InvoiceForm,
            extraProps: { editable: false },
            actions: [
              {
                title: 'back',
                action: ({ blockId }) => eifActions.popLayoutBlock(blockId),
              }, {
                title: 'download',
                action: (block) => {
                  const invoice = eifActions.getDataAtPath<InvoiceFormData>(block.dataPath);
                  const { month, year } = eifActions.getData(state => state.selectedMonth);
                  const data = getDateFrom(1, month, year);
                  // @ts-ignore
                  window.html2pdf(document.getElementById('to-pdf'), {
                    filename: `Factura ${invoice?.name} - ${data.format('MMM YYYY')}.pdf`,
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
              }
            ],
          });
        }
      },
    ],
  });
}
