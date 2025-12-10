import { Account, CreateAccountRequest, UpdateAccountRequest } from "@/types";
import { apiClient } from "./client";

export const accountsApi = {
  getAll: async (): Promise<Account[]> => {
    return apiClient.get<Account[]>("/accounts");
  },

  getByAccountNo: async (accountNo: number): Promise<Account> => {
    return apiClient.get<Account>(`/accounts/${accountNo}`);
  },

  getByGroup: async (group: number): Promise<Account[]> => {
    return apiClient.get<Account[]>(`/accounts/group/${group}`);
  },

  create: async (data: CreateAccountRequest): Promise<Account> => {
    return apiClient.post<Account>("/accounts", data);
  },

  update: async (accountNo: number, data: UpdateAccountRequest): Promise<Account> => {
    return apiClient.put<Account>(`/accounts/${accountNo}`, data);
  },

  delete: async (accountNo: number): Promise<void> => {
    return apiClient.delete<void>(`/accounts/${accountNo}`);
  },
};
