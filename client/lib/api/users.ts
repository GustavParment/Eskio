import { User, CreateUserRequest, UpdateUserRequest } from "@/types";
import { apiClient } from "./client";

export const usersApi = {
  getById: async (id: number): Promise<User> => {
    return apiClient.get<User>(`/users/${id}`);
  },

  getByEmail: async (email: string): Promise<User> => {
    return apiClient.get<User>(`/users/email/${email}`);
  },

  create: async (data: CreateUserRequest): Promise<User> => {
    return apiClient.post<User>("/users", data);
  },

  update: async (id: number, data: UpdateUserRequest): Promise<User> => {
    return apiClient.put<User>(`/users/${id}`, data);
  },

  delete: async (id: number): Promise<void> => {
    return apiClient.delete<void>(`/users/${id}`);
  },
};
