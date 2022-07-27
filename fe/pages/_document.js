import { Head, Html, NextScript, Main } from "next/document";

export default function Document() {
  return (
    <Html lang="en">
      <Head>
        <link rel="preconnect" href="https://fonts.googleapis.com" />
        <link
          rel="preconnect"
          href="https://fonts.gstatic.com"
          crossOrigin={true}
        />
        <link
          href="https://fonts.googleapis.com/css2?family=Inter:wght@300;400;600;800&display=swap"
          rel="stylesheet"
        ></link>
      </Head>
      <body className="tracking-tight overflow-x-hidden font-inter text-gray-700">
        <Main />
        <NextScript />
      </body>
    </Html>
  );
}
