"use client";

import DashboardLayout from "@/components/layout/DashboardLayout";

export default function ReportsPage() {
  const reportTypes = [
    {
      name: "Resultaträkning",
      description: "Visa företagets intäkter och kostnader",
      icon: "📊",
      color: "bg-blue-500",
    },
    {
      name: "Balansräkning",
      description: "Visa företagets tillgångar och skulder",
      icon: "⚖️",
      color: "bg-green-500",
    },
    {
      name: "Kontoutdrag",
      description: "Visa transaktioner för specifika konton",
      icon: "📋",
      color: "bg-purple-500",
    },
    {
      name: "Huvudbok",
      description: "Visa alla bokförda verifikat",
      icon: "📚",
      color: "bg-orange-500",
    },
    {
      name: "Momsrapport",
      description: "Visa momsunderlag och beräkningar",
      icon: "💶",
      color: "bg-red-500",
    },
    {
      name: "Period jämförelse",
      description: "Jämför resultat mellan olika perioder",
      icon: "📈",
      color: "bg-indigo-500",
    },
  ];

  return (
    <DashboardLayout>
      {/* Header */}
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-900">Rapporter</h1>
        <p className="text-gray-600 mt-2">Visa och exportera finansiella rapporter</p>
      </div>

      {/* Reports grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {reportTypes.map((report) => (
          <button
            key={report.name}
            className="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-shadow text-left"
          >
            <div className={`${report.color} w-12 h-12 rounded-lg flex items-center justify-center text-2xl mb-4`}>
              {report.icon}
            </div>
            <h3 className="text-lg font-semibold text-gray-900 mb-2">{report.name}</h3>
            <p className="text-sm text-gray-600">{report.description}</p>
            <div className="mt-4 text-sm text-blue-600 font-medium">
              Visa rapport →
            </div>
          </button>
        ))}
      </div>

      {/* Coming soon notice */}
      <div className="mt-12 bg-blue-50 border border-blue-200 rounded-xl p-6">
        <h3 className="text-lg font-semibold text-blue-900 mb-2">🚧 Under utveckling</h3>
        <p className="text-blue-800">
          Rapportfunktionen är under utveckling. Snart kommer du att kunna generera och exportera
          alla typer av finansiella rapporter direkt från systemet.
        </p>
      </div>
    </DashboardLayout>
  );
}
