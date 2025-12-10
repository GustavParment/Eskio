import { LineItem, CreateLineItemRequest, UpdateLineItemRequest } from "@/types";
import { apiClient } from "./client";

export const lineItemsApi = {
  getById: async (id: number): Promise<LineItem> => {
    return apiClient.get<LineItem>(`/lineitems/${id}`);
  },

  getByVoucher: async (voucherId: number): Promise<LineItem[]> => {
    return apiClient.get<LineItem[]>(`/lineitems/voucher/${voucherId}`);
  },

  getByAccount: async (accountNo: number): Promise<LineItem[]> => {
    return apiClient.get<LineItem[]>(`/lineitems/account/${accountNo}`);
  },

  create: async (data: CreateLineItemRequest): Promise<LineItem> => {
    return apiClient.post<LineItem>("/lineitems", data);
  },

  update: async (id: number, data: UpdateLineItemRequest): Promise<LineItem> => {
    return apiClient.put<LineItem>(`/lineitems/${id}`, data);
  },

  delete: async (id: number): Promise<void> => {
    return apiClient.delete<void>(`/lineitems/${id}`);
  },
};
