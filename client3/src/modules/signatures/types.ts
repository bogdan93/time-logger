export interface Signature {
  id: string;
  name: string;
  field?: string;
}

export interface SignaturesMap {
  [id: string]: Signature;
}
