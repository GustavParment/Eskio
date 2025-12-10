"use client";

import { useState, useRef, useEffect } from "react";
import { Account } from "@/types";

interface AccountSearchProps {
  accounts: Account[];
  value: string;
  onChange: (accountNo: string) => void;
  placeholder?: string;
}

export default function AccountSearch({
  accounts,
  value,
  onChange,
  placeholder = "Sök konto...",
}: AccountSearchProps) {
  const [isOpen, setIsOpen] = useState(false);
  const [searchTerm, setSearchTerm] = useState("");
  const [highlightedIndex, setHighlightedIndex] = useState(0);
  const inputRef = useRef<HTMLInputElement>(null);
  const listRef = useRef<HTMLUListElement>(null);
  const containerRef = useRef<HTMLDivElement>(null);

  // Find the selected account
  const selectedAccount = accounts.find(
    (acc) => acc.account_no.toString() === value
  );

  // Filter accounts based on search term
  const filteredAccounts = accounts.filter((account) => {
    if (!searchTerm) return true;
    const search = searchTerm.toLowerCase();
    return (
      account.account_no.toString().includes(search) ||
      account.account_name.toLowerCase().includes(search)
    );
  });

  // Reset highlighted index when filtered results change
  useEffect(() => {
    setHighlightedIndex(0);
  }, [searchTerm]);

  // Scroll highlighted item into view
  useEffect(() => {
    if (isOpen && listRef.current) {
      const highlightedElement = listRef.current.children[highlightedIndex] as HTMLElement;
      if (highlightedElement) {
        highlightedElement.scrollIntoView({ block: "nearest" });
      }
    }
  }, [highlightedIndex, isOpen]);

  // Close dropdown when clicking outside
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (containerRef.current && !containerRef.current.contains(event.target as Node)) {
        setIsOpen(false);
        setSearchTerm("");
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  const handleSelect = (account: Account) => {
    onChange(account.account_no.toString());
    setIsOpen(false);
    setSearchTerm("");
  };

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (!isOpen) {
      if (e.key === "Enter" || e.key === "ArrowDown" || e.key === " ") {
        e.preventDefault();
        setIsOpen(true);
      }
      return;
    }

    switch (e.key) {
      case "ArrowDown":
        e.preventDefault();
        setHighlightedIndex((prev) =>
          prev < filteredAccounts.length - 1 ? prev + 1 : prev
        );
        break;
      case "ArrowUp":
        e.preventDefault();
        setHighlightedIndex((prev) => (prev > 0 ? prev - 1 : 0));
        break;
      case "Enter":
        e.preventDefault();
        if (filteredAccounts[highlightedIndex]) {
          handleSelect(filteredAccounts[highlightedIndex]);
        }
        break;
      case "Escape":
        e.preventDefault();
        setIsOpen(false);
        setSearchTerm("");
        break;
      case "Tab":
        setIsOpen(false);
        setSearchTerm("");
        break;
    }
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(e.target.value);
    if (!isOpen) {
      setIsOpen(true);
    }
  };

  const handleFocus = () => {
    setIsOpen(true);
  };

  return (
    <div ref={containerRef} className="relative">
      <div
        className={`flex items-center border rounded cursor-text ${
          isOpen
            ? "border-blue-500 ring-2 ring-blue-500"
            : "border-gray-300 hover:border-gray-400"
        }`}
        onClick={() => inputRef.current?.focus()}
      >
        <input
          ref={inputRef}
          type="text"
          value={isOpen ? searchTerm : selectedAccount ? `${selectedAccount.account_no} - ${selectedAccount.account_name}` : ""}
          onChange={handleInputChange}
          onFocus={handleFocus}
          onKeyDown={handleKeyDown}
          placeholder={selectedAccount ? "" : placeholder}
          className="w-full px-2 py-2 text-sm text-black bg-transparent outline-none"
        />
        <div className="px-2 text-gray-400">
          <svg
            className={`w-4 h-4 transition-transform ${isOpen ? "rotate-180" : ""}`}
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
          </svg>
        </div>
      </div>

      {isOpen && (
        <ul
          ref={listRef}
          className="absolute z-50 w-full mt-1 bg-white border border-gray-300 rounded-lg shadow-lg max-h-60 overflow-auto"
        >
          {filteredAccounts.length === 0 ? (
            <li className="px-3 py-2 text-sm text-gray-500">
              Inga konton hittades
            </li>
          ) : (
            filteredAccounts.map((account, index) => (
              <li
                key={account.account_no}
                onClick={() => handleSelect(account)}
                className={`px-3 py-2 text-sm cursor-pointer ${
                  index === highlightedIndex
                    ? "bg-blue-50 text-blue-900"
                    : "text-gray-900 hover:bg-gray-50"
                } ${value === account.account_no.toString() ? "font-medium" : ""}`}
              >
                <span className="font-mono text-blue-600">{account.account_no}</span>
                <span className="mx-2 text-gray-400">-</span>
                <span>{account.account_name}</span>
              </li>
            ))
          )}
        </ul>
      )}
    </div>
  );
}
