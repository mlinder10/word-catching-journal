import type { Definition } from "@/types";
import { Label } from "./ui/label";
import { Skeleton } from "./ui/skeleton";
import { capitalize } from "@/lib/utils";

type DefinitionViewProps = {
  definition: Definition;
  full?: boolean;
};

export default function DefinitionView({
  definition,
  full = true,
}: DefinitionViewProps) {
  return (
    <div className="space-y-2">
      <div className="flex flex-wrap items-center gap-2">
        <span className="font-semibold">{capitalize(definition.word)}</span>
        {definition.pronunciation && <span>({definition.pronunciation})</span>}
        <span className="text-muted-foreground italic">
          {capitalize(definition.partOfSpeech)}
        </span>
      </div>

      <div>
        <Label className="text-muted-foreground text-sm">Definition</Label>
        <p className="ml-2">{definition.definition}</p>
      </div>

      {full && (
        <div>
          <Label className="text-muted-foreground text-sm">Example</Label>
          <p className="ml-2">{definition.example}</p>
        </div>
      )}

      {full && definition.synonyms.length > 0 && (
        <div>
          <Label className="text-muted-foreground text-sm">Synonyms</Label>
          <p className="ml-2">
            {definition.synonyms.map(capitalize).join(", ")}
          </p>
        </div>
      )}

      {full && definition.antonyms.length > 0 && (
        <div>
          <Label className="text-muted-foreground text-sm">Antonyms</Label>
          <p className="ml-2">
            {definition.antonyms.map(capitalize).join(", ")}
          </p>
        </div>
      )}
    </div>
  );
}

export function DefinitionViewSkeleton() {
  return (
    <div className="space-y-2">
      <div className="flex gap-2">
        <Skeleton className="flex-2 rounded-md h-6" />
        <Skeleton className="flex-2 rounded-md h-6" />
        <span className="flex-3" />
      </div>
      <Skeleton className="rounded-md w-full h-24" />
      <Skeleton className="rounded-full w-full h-6" />
    </div>
  );
}
