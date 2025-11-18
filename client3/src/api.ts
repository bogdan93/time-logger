import PocketBase from 'pocketbase';

// @ts-ignore
const APP_API = import.meta.env.VITE_APP_URL;
const api = new PocketBase(APP_API);

export function getImageLink(record: any, fileName: string) {
  return api.records.getFileUrl(record, fileName);
}

export default api;
