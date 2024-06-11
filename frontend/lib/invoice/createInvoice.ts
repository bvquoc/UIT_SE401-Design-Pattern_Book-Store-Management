import { endPoint } from "@/constants";
import axios from "axios";
import { getApiKey } from "../auth/action";

type InvoiceProps = {
  customerId: string;
  isUsePoint: boolean;
  details: {
    bookId: string;
    qty: number;
  }[];
};
export default async function createInvoice(data: InvoiceProps) {
  const url = `${endPoint}/v1/invoices`;

  const token = await getApiKey();
  const headers = {
    accept: "application/json",
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,
    // Add other headers as needed
  };

  // Make a POST request with headers
  const res = axios
    .post(url, data, { headers: headers })
    .then((response) => {
      if (response) return response.data;
    })
    .catch((error) => {
      console.error("Error:", error);
      return error.response.data;
    });
  return res;
}
