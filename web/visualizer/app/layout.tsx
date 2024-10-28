import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "CodeGraph Visualizer",
  description: "To visualize code graph",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        {children}
      </body>
    </html>
  );
}
