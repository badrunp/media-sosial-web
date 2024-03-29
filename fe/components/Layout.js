import Navbar from "./Navbar";

export default function Layout({ children }) {
  return (
    <>
      <Navbar />
      <div className="mt-10">{children}</div>
    </>
  );
}
