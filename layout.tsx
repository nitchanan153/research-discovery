import type { Metadata } from "next";
import "./globals.css";
export const metadata: Metadata = { title: "Bio-RD Intelligence", description: "Research discovery workspace" };
export default function RootLayout({ children }: Readonly<{ children: React.ReactNode }>) { return <html lang="th"><body>{children}</body></html>; }
