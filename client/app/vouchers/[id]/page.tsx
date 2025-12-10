"use client";

import { useEffect, useState } from "react";
import { useParams, useRouter } from "next/navigation";
import DashboardLayout from "@/components/layout/DashboardLayout";
import { vouchersApi } from "@/lib/api/vouchers";
import { lineItemsApi } from "@/lib/api/lineitems";
import { accountsApi } from "@/lib/api/accounts";
import { Voucher, LineItem, Account } from "@/types";
import Link from "next/link";

export default function VoucherDetailPage() {
  const params = useParams();
  const router = useRouter();
  const voucherId = Number(params.id);

  const [voucher, setVoucher] = useState<Voucher | null>(null);
  const [lineItems, setLineItems] = useState<LineItem[]>([]);
  const [accounts, setAccounts] = useState<Record<number, Account>>({});
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchData = async () => {
      try {
        const [voucherData, lineItemsData, accountsData] = await Promise.all([
          vouchersApi.getById(voucherId),
          lineItemsApi.getByVoucher(voucherId),
          accountsApi.getAll(),
        ]);

        setVoucher(voucherData);
        setLineItems(Array.isArray(lineItemsData) ? lineItemsData : []);

        // Create account lookup map
        const accountMap: Record<number, Account> = {};
        (accountsData || []).forEach((acc: Account) => {
          accountMap[acc.account_no] = acc;
        });
        setAccounts(accountMap);
      } catch (err) {
        setError(err instanceof Error ? err.message : "Failed to load voucher");
      } finally {
        setLoading(false);
      }
    };

    if (voucherId) {
      fetchData();
    }
  }, [voucherId]);

  const totalDebit = lineItems.reduce((sum, item) => sum + (item.debit_amount || 0), 0);
  const totalCredit = lineItems.reduce((sum, item) => sum + (item.credit_amount || 0), 0);
  const isBalanced = Math.abs(totalDebit - totalCredit) < 0.01;

  if (loading) {
    return (
      <DashboardLayout>
        <div className="flex items-center justify-center h-64">
          <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
        </div>
      </DashboardLayout>
    );
  }

  if (error || !voucher) {
    return (
      <DashboardLayout>
        <div className="text-center py-12">
          <p className="text-red-600 mb-4">{error || "Verifikatet hittades inte"}</p>
          <Link
            href="/vouchers"
            className="text-blue-600 hover:text-blue-700 font-medium"
          >
            Tillbaka till verifikat
          </Link>
        </div>
      </DashboardLayout>
    );
  }

  return (
    <DashboardLayout>
      <div className="max-w-5xl mx-auto">
        {/* Header */}
        <div className="mb-8">
          <button
            onClick={() => router.back()}
            className="text-sm text-gray-600 hover:text-gray-900 mb-4 flex items-center gap-2"
          >
            &larr; Tillbaka
          </button>
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-3xl font-bold text-gray-900">
                Verifikat #{voucher.voucher_id}
              </h1>
              <p className="text-gray-600 mt-2">{voucher.description}</p>
            </div>
            <div className="flex gap-3">
              <Link
                href={`/vouchers/${voucher.voucher_id}/edit`}
                className="px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors font-medium"
              >
                Redigera
              </Link>
            </div>
          </div>
        </div>

        {/* Voucher Info Card */}
        <div className="bg-white rounded-xl shadow-sm border border-gray-200 p-6 mb-6">
          <h2 className="text-lg font-semibold text-gray-900 mb-4">Verifikatuppgifter</h2>
          <div className="grid grid-cols-2 md:grid-cols-4 gap-6">
            <div>
              <p className="text-sm text-gray-500 mb-1">Datum</p>
              <p className="font-medium text-gray-900">
                {new Date(voucher.date).toLocaleDateString("sv-SE")}
              </p>
            </div>
            <div>
              <p className="text-sm text-gray-500 mb-1">Period</p>
              <p className="font-medium text-gray-900">
                {new Date(voucher.period + "-01").toLocaleDateString("sv-SE", {
                  year: "numeric",
                  month: "long",
                })}
              </p>
            </div>
            <div>
              <p className="text-sm text-gray-500 mb-1">Referens</p>
              <p className="font-medium text-gray-900">{voucher.reference || "-"}</p>
            </div>
            <div>
              <p className="text-sm text-gray-500 mb-1">Totalbelopp</p>
              <p className="font-medium text-gray-900">
                {voucher.total_amount.toLocaleString("sv-SE", {
                  minimumFractionDigits: 2,
                  maximumFractionDigits: 2,
                })}{" "}
                kr
              </p>
            </div>
          </div>

          <div className="mt-4 pt-4 border-t border-gray-100">
            <p className="text-sm text-gray-500 mb-1">Beskrivning</p>
            <p className="text-gray-900">{voucher.description}</p>
          </div>

          <div className="mt-4 pt-4 border-t border-gray-100 flex gap-8 text-sm text-gray-500">
            <div>
              Skapad:{" "}
              {new Date(voucher.created_at).toLocaleString("sv-SE")}
            </div>
            <div>
              Uppdaterad:{" "}
              {new Date(voucher.updated_at).toLocaleString("sv-SE")}
            </div>
          </div>
        </div>

        {/* Line Items Card */}
        <div className="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden mb-6">
          <div className="p-6 border-b border-gray-200">
            <h2 className="text-lg font-semibold text-gray-900">Konteringsrader</h2>
          </div>

          {lineItems.length === 0 ? (
            <div className="text-center py-12">
              <p className="text-gray-500">Inga konteringsrader</p>
            </div>
          ) : (
            <div className="overflow-x-auto">
              <table className="w-full">
                <thead className="bg-gray-50">
                  <tr>
                    <th className="text-left py-3 px-6 text-sm font-semibold text-gray-700">
                      Konto
                    </th>
                    <th className="text-left py-3 px-6 text-sm font-semibold text-gray-700">
                      Kontonamn
                    </th>
                    <th className="text-right py-3 px-6 text-sm font-semibold text-gray-700">
                      Debet
                    </th>
                    <th className="text-right py-3 px-6 text-sm font-semibold text-gray-700">
                      Kredit
                    </th>
                    <th className="text-center py-3 px-6 text-sm font-semibold text-gray-700">
                      Moms
                    </th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-gray-100">
                  {lineItems.map((item) => (
                    <tr key={item.line_id} className="hover:bg-gray-50">
                      <td className="py-4 px-6 text-sm font-medium text-gray-900">
                        {item.account_no}
                      </td>
                      <td className="py-4 px-6 text-sm text-gray-600">
                        {accounts[item.account_no]?.account_name || "-"}
                      </td>
                      <td className="py-4 px-6 text-sm text-gray-900 text-right">
                        {item.debit_amount > 0
                          ? item.debit_amount.toLocaleString("sv-SE", {
                              minimumFractionDigits: 2,
                              maximumFractionDigits: 2,
                            })
                          : "-"}
                      </td>
                      <td className="py-4 px-6 text-sm text-gray-900 text-right">
                        {item.credit_amount > 0
                          ? item.credit_amount.toLocaleString("sv-SE", {
                              minimumFractionDigits: 2,
                              maximumFractionDigits: 2,
                            })
                          : "-"}
                      </td>
                      <td className="py-4 px-6 text-sm text-gray-600 text-center">
                        {item.tax_code}%
                      </td>
                    </tr>
                  ))}
                </tbody>
                <tfoot className="bg-gray-50">
                  <tr className="font-semibold">
                    <td colSpan={2} className="py-4 px-6 text-sm text-gray-900">
                      Summa
                    </td>
                    <td className="py-4 px-6 text-sm text-gray-900 text-right">
                      {totalDebit.toLocaleString("sv-SE", {
                        minimumFractionDigits: 2,
                        maximumFractionDigits: 2,
                      })}
                    </td>
                    <td className="py-4 px-6 text-sm text-gray-900 text-right">
                      {totalCredit.toLocaleString("sv-SE", {
                        minimumFractionDigits: 2,
                        maximumFractionDigits: 2,
                      })}
                    </td>
                    <td></td>
                  </tr>
                </tfoot>
              </table>
            </div>
          )}
        </div>

        {/* Balance Status */}
        <div
          className={`p-4 rounded-lg border ${
            isBalanced
              ? "bg-green-50 border-green-200"
              : "bg-red-50 border-red-200"
          }`}
        >
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-2">
              <span
                className={`text-lg ${
                  isBalanced ? "text-green-600" : "text-red-600"
                }`}
              >
                {isBalanced ? "Balanserat" : "Ej balanserat"}
              </span>
            </div>
            <div className="text-sm">
              <span className="text-gray-600">Differens: </span>
              <span
                className={`font-medium ${
                  isBalanced ? "text-green-600" : "text-red-600"
                }`}
              >
                {Math.abs(totalDebit - totalCredit).toLocaleString("sv-SE", {
                  minimumFractionDigits: 2,
                  maximumFractionDigits: 2,
                })}{" "}
                kr
              </span>
            </div>
          </div>
        </div>
      </div>
    </DashboardLayout>
  );
}
