import { apiClient } from "./client";

export interface IncomeStatementEntry {
  account_no: number;
  account_name: string;
  balance: number;
}

export interface IncomeStatement {
  period: {
    from_date: string;
    to_date: string;
  };
  income: IncomeStatementEntry[];
  expenses: IncomeStatementEntry[];
  total_income: number;
  total_expenses: number;
  net_result: number;
}

export const reportsApi = {
  getIncomeStatement: async (fromDate: string, toDate: string): Promise<IncomeStatement> => {
    return apiClient.get<IncomeStatement>(`/reports/income-statement?from_date=${fromDate}&to_date=${toDate}`);
  },
};
