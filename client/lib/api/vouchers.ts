import { Voucher, CreateVoucherRequest, UpdateVoucherRequest, ValidationResponse } from "@/types";
import { apiClient } from "./client";

export const vouchersApi = {
  getAll: async (): Promise<Voucher[]> => {
    return apiClient.get<Voucher[]>("/vouchers");
  },

  getById: async (id: number): Promise<Voucher> => {
    return apiClient.get<Voucher>(`/vouchers/${id}`);
  },

  getByPeriod: async (period: string): Promise<Voucher[]> => {
    return apiClient.get<Voucher[]>(`/vouchers/period/${period}`);
  },

  getByUser: async (userId: number): Promise<Voucher[]> => {
    return apiClient.get<Voucher[]>(`/vouchers/user/${userId}`);
  },

  validate: async (id: number): Promise<ValidationResponse> => {
    return apiClient.get<ValidationResponse>(`/vouchers/${id}/validate`);
  },

  create: async (data: CreateVoucherRequest): Promise<Voucher> => {
    return apiClient.post<Voucher>("/vouchers", data);
  },

  update: async (id: number, data: UpdateVoucherRequest): Promise<Voucher> => {
    return apiClient.put<Voucher>(`/vouchers/${id}`, data);
  },

  delete: async (id: number): Promise<void> => {
    return apiClient.delete<void>(`/vouchers/${id}`);
  },
};
