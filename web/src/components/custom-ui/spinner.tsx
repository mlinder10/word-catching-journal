import { Loader } from "lucide-react";

export default function Spinner() {
  return (
    <div className="flex justify-center items-center">
      <Loader className="animate-spin" />
    </div>
  );
}
