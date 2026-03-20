import PageTitle from "@/components/custom-ui/page-title";
import { createFileRoute } from "@tanstack/react-router";
import z from "zod";
import DefinitionsGrid from "@/components/pages/define/definitions-grid";
import DefinitionInput from "@/components/pages/define/definition-input";

export const Route = createFileRoute("/app/_authenticated/define")({
  validateSearch: z.object({
    word: z.string().optional(),
  }),
  component: DefinePage,
});

function DefinePage() {
  return (
    <main className="flex flex-col gap-4 h-full">
      <PageTitle title="Define" />
      <DefinitionInput />
      <DefinitionsGrid />
    </main>
  );
}
