import { useNavigate, useSearch } from "@tanstack/react-router";
import { DefinitionViewSkeleton } from "@/components/definition-view";
import { Card, CardContent } from "@/components/ui/card";
import DefinitionItem from "./definition-item";
import api from "@/api";
import { useQuery } from "@tanstack/react-query";
import {
  Empty,
  EmptyDescription,
  EmptyHeader,
  EmptyMedia,
  EmptyTitle,
} from "@/components/ui/empty";
import { Search } from "lucide-react";
import { Separator } from "@/components/ui/separator";

function useDefineWord(word: string) {
  return useQuery({
    queryKey: ["define", word],
    queryFn: async () => api.defineWord(word),
    enabled: !!word,
  });
}

export default function DefinitionsGrid() {
  const { word } = useSearch({ from: "/app/_authenticated/define" });
  const { data, isLoading, error } = useDefineWord(word ?? "");

  if (isLoading) return <DefinitionsGridSkeleton />;
  if (error) return <DefinitionsGridError error={error} />;
  if (!data) return <EmptyDefinitionsGrid />;
  if (data.type === "suggestions")
    return <SuggestionsView words={data.suggestions} />;

  return (
    <section className="flex-1 px-4 pb-4 overflow-y-auto">
      <ul className="gap-4 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
        {data.definitions.map((d, i) => (
          <DefinitionItem key={i + d.definition} definition={d} />
        ))}
      </ul>
    </section>
  );
}

function DefinitionsGridSkeleton() {
  return (
    <section className="flex-1 px-4 pb-4 overflow-y-auto">
      <ul className="gap-4 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
        {Array.from({ length: 9 }).map((_, i) => (
          <li key={i}>
            <Card>
              <CardContent>
                <DefinitionViewSkeleton />
              </CardContent>
            </Card>
          </li>
        ))}
      </ul>
    </section>
  );
}

function DefinitionsGridError({ error }: { error: Error }) {
  return (
    <section className="flex-1 place-items-center grid">
      <p className="text-red-500 text-sm">{error.message}</p>
    </section>
  );
}

function EmptyDefinitionsGrid() {
  return (
    <section className="flex-1 place-items-center grid">
      <Empty>
        <EmptyHeader>
          <EmptyMedia variant="icon">
            <Search />
          </EmptyMedia>
          <EmptyTitle>Search for a word&apos;s meaning</EmptyTitle>
          <EmptyDescription>
            Enter a word in the search bar above and click &quot;Define&quot; to
            get started
          </EmptyDescription>
        </EmptyHeader>
      </Empty>
    </section>
  );
}

type SuggestionsProps = {
  words: string[];
};

function SuggestionsView({ words }: SuggestionsProps) {
  const navigate = useNavigate();

  function handleSelect(word: string) {
    navigate({
      to: "/app/define",
      search: { word },
    });
  }

  return (
    <section className="flex-1 px-4 pb-4 overflow-y-auto">
      <Card className="py-0">
        <ul>
          {words.map((w, i) => (
            <li key={i} onClick={() => handleSelect(w)} className="px-4">
              <p className="py-2 text-base cursor-pointer">{w}</p>
              {i !== words.length - 1 && <Separator />}
            </li>
          ))}
        </ul>
      </Card>
    </section>
  );
}
