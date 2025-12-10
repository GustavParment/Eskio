# Bookkeeper Frontend

Modern, secure Next.js frontend for the Bookkeeper Swedish accounting application.

## 🚀 Tech Stack

- **Next.js 15** (App Router)
- **React 19**
- **TypeScript 5**
- **Tailwind CSS 4**
- Secure httpOnly cookie authentication

## 📁 Project Structure

```
client/app/
├── app/                          # Next.js app directory
│   ├── auth/
│   │   ├── login/                # Login page
│   │   └── register/             # Registration page
│   ├── dashboard/                # Dashboard overview
│   ├── accounts/                 # Account management
│   ├── vouchers/                 # Voucher management
│   ├── reports/                  # Financial reports
│   ├── layout.tsx                # Root layout with AuthProvider
│   └── page.tsx                  # Home (redirects to login)
│
├── components/
│   └── layout/
│       ├── Sidebar.tsx           # Navigation sidebar
│       ├── DashboardLayout.tsx   # Main app layout
│       └── ProtectedRoute.tsx    # Auth guard component
│
├── lib/
│   ├── api/
│   │   ├── client.ts             # Base API client with httpOnly cookies
│   │   ├── auth.ts               # Authentication API
│   │   ├── accounts.ts           # Accounts API
│   │   ├── vouchers.ts           # Vouchers API
│   │   ├── lineitems.ts          # Line items API
│   │   └── users.ts              # Users API
│   ├── contexts/
│   │   └── AuthContext.tsx       # Authentication context provider
│   └── utils/                    # Utility functions (future)
│
└── types/
    └── index.ts                  # TypeScript type definitions
```

## 🔐 Security Features

### httpOnly Cookies (Secure Authentication)

We use **httpOnly cookies** instead of localStorage for JWT tokens to prevent XSS attacks:

- Cookies are set by the backend with `httpOnly` flag
- Frontend automatically sends cookies with `credentials: "include"`
- JavaScript cannot access the tokens (XSS protection)
- CSRF protection via SameSite cookie attribute

### CORS Configuration Required

Your Go backend needs CORS with credentials support:

```go
// Add to your backend
import "github.com/gin-contrib/cors"

router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3000"},
    AllowCredentials: true,
    AllowHeaders:     []string{"Content-Type", "Authorization"},
}))
```

## ⚠️ Backend Changes Required

To use this frontend, your Go backend needs these modifications:

### 1. Update Authentication Endpoints

**Current:** Returns JWT in response body
**Needed:** Set JWT as httpOnly cookie

```go
// Login handler example
func (h *AuthHandler) Login(c *gin.Context) {
    // ... authentication logic ...

    // Instead of returning token in body:
    // c.JSON(200, gin.H{"token": token, "user": user})

    // Set as httpOnly cookie:
    c.SetCookie(
        "token",           // name
        token,             // value
        604800,            // maxAge (7 days)
        "/",               // path
        "",                // domain
        false,             // secure (set true in production with HTTPS)
        true,              // httpOnly
    )

    c.JSON(200, gin.H{"user": user})
}
```

### 2. Add Logout Endpoint

```go
func (h *AuthHandler) Logout(c *gin.Context) {
    c.SetCookie("token", "", -1, "/", "", false, true)
    c.JSON(200, gin.H{"message": "Logged out successfully"})
}

// Add route
authRoutes.POST("/logout", authHandler.Logout)
```

### 3. Update JWT Middleware

Read token from cookie instead of Authorization header:

```go
func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get token from cookie
        token, err := c.Cookie("token")
        if err != nil {
            c.JSON(401, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // ... validate token ...
    }
}
```

### 4. Environment Variables

Add to your backend `.env`:

```
ALLOWED_ORIGINS=http://localhost:3000
COOKIE_DOMAIN=localhost
COOKIE_SECURE=false  # Set to true in production
```

## 🎨 Features Implemented

### ✅ Completed

- **Authentication**
  - Login page with Swedish UI
  - Registration page with role selection
  - Secure httpOnly cookie authentication
  - Protected routes with loading states
  - Auth context with user state management

- **Dashboard**
  - Overview stats (vouchers, accounts, period)
  - Recent vouchers table
  - Quick action cards
  - Period-based filtering

- **Account Management**
  - Account list with BAS group filtering
  - Search by account number or name
  - Group filter (1-8)
  - Responsive table design

- **Voucher Management**
  - Voucher list with period filtering
  - Search functionality
  - Swedish locale formatting
  - Period selector (last 12 months)

- **Navigation**
  - Sidebar with icons
  - Active route highlighting
  - User profile section
  - Admin-only menu items
  - Logout button

- **Reports**
  - Placeholder page with report types
  - Coming soon notice

### 🚧 To Be Implemented

- **Account Detail Pages**
  - View account transactions
  - Edit account form
  - Delete account

- **Voucher Creation/Editing**
  - Create voucher form with line items
  - Edit existing vouchers
  - Balance validation indicator
  - Double-entry accounting validation

- **User Management** (Admin only)
  - User list
  - Create/edit users
  - Role management

- **Reports**
  - Income statement (Resultaträkning)
  - Balance sheet (Balansräkning)
  - Account ledger
  - VAT reports
  - PDF export

- **Additional Features**
  - Toast notifications
  - Error boundaries
  - Loading skeletons
  - Pagination
  - Sorting
  - Date pickers
  - Modal dialogs

## 🚀 Getting Started

### Prerequisites

- Node.js 18+
- Running Go backend on `http://localhost:8080`

### Installation

```bash
cd client/app
npm install
```

### Environment Setup

Create `.env.local`:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
```

### Development

```bash
npm run dev
```

Open http://localhost:3000

### Build

```bash
npm run build
npm start
```

## 📝 API Integration

All API calls use the centralized `apiClient` with:

- Automatic cookie handling
- Type-safe requests/responses
- Error handling
- TypeScript interfaces matching backend models

Example usage:

```typescript
import { accountsApi } from '@/lib/api/accounts';

// Fetch all accounts
const accounts = await accountsApi.getAll();

// Create account
const newAccount = await accountsApi.create({
  account_no: 1930,
  account_name: "Bank Account",
  account_group: 1,
  tax_standard: "0%",
  type: "BS",
  standard_side: "Debit"
});
```

## 🎨 Design System

### Colors

- Primary: Blue (`bg-blue-600`)
- Success: Green (`bg-green-500`)
- Warning: Orange (`bg-orange-500`)
- Error: Red (`bg-red-500`)
- Gray scale for text and backgrounds

### Typography

- Headings: `font-bold`
- Body: `font-medium` or default
- Monospace numbers for accounting values

### Components

- Rounded corners: `rounded-lg` or `rounded-xl`
- Shadows: `shadow-sm` for cards
- Hover states on all interactive elements
- Focus rings: `focus:ring-2 focus:ring-blue-500`

## 🌍 Localization

Currently using Swedish:

- All UI text in Swedish
- Swedish date formatting (`sv-SE`)
- Swedish number formatting with spaces
- Month names in Swedish

## 🧪 Testing Checklist

Before deploying, test:

- [ ] Login with valid credentials
- [ ] Login with invalid credentials
- [ ] Registration flow
- [ ] Logout functionality
- [ ] Protected route redirection
- [ ] Dashboard loads data
- [ ] Account filtering and search
- [ ] Voucher filtering by period
- [ ] Period dropdown selection
- [ ] Navigation between pages
- [ ] Responsive design (mobile/tablet)
- [ ] Loading states
- [ ] Error messages

## 🔒 Security Considerations

1. **httpOnly Cookies:** Tokens cannot be accessed via JavaScript
2. **CORS:** Backend must allow credentials from frontend origin
3. **CSRF:** Use SameSite cookie attribute
4. **HTTPS:** Required in production for secure cookies
5. **Input Validation:** Validate all user inputs
6. **SQL Injection:** Backend uses parameterized queries
7. **XSS Protection:** React escapes output by default

## 📚 Resources

- [Next.js Documentation](https://nextjs.org/docs)
- [Tailwind CSS](https://tailwindcss.com/docs)
- [BAS Account System](https://www.bas.se/)
- [Swedish Bookkeeping Standards](https://www.bokforingsnamnden.se/)

## 🤝 Contributing

When adding new features:

1. Add TypeScript types to `types/index.ts`
2. Create API functions in `lib/api/`
3. Build components in `components/`
4. Add pages in `app/`
5. Use consistent styling with Tailwind
6. Follow Swedish naming for accounting terms

## 📄 License

[Your License Here]
