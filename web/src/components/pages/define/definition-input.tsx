import { isAlpha } from "@/lib/utils";
import { useNavigate, useSearch } from "@tanstack/react-router";
import { useState } from "react";
import {
  InputGroup,
  InputGroupAddon,
  InputGroupButton,
  InputGroupInput,
} from "@/components/ui/input-group";
import { Search } from "lucide-react";

export default function DefinitionInput() {
  const { word: initialWord } = useSearch({
    from: "/app/_authenticated/define",
  });
  const [word, setWord] = useState(initialWord ?? "");
  const [inputError, setInputError] = useState<string | null>(null);
  const navigate = useNavigate();

  function handleDefine() {
    const trimmed = word.trim();
    if (!trimmed) return setInputError("Word is required");
    if (!isAlpha(trimmed))
      return setInputError("Word must only contain letters and hyphens");
    setInputError(null);
    navigate({ to: "/app/define", search: { word } });
  }

  function handleKeydown(e: React.KeyboardEvent<HTMLInputElement>) {
    if (e.key === "Enter") {
      e.preventDefault();
      handleDefine();
    }
  }

  return (
    <section className="space-y-2 px-4">
      <InputGroup>
        <InputGroupInput
          placeholder="Define a word..."
          value={word}
          onChange={(e) => setWord(e.target.value)}
          onKeyDown={handleKeydown}
        />
        <InputGroupAddon align="inline-start">
          <Search />
        </InputGroupAddon>
        <InputGroupAddon align="inline-end">
          <InputGroupButton variant="default" onClick={handleDefine}>
            Define
          </InputGroupButton>
        </InputGroupAddon>
      </InputGroup>
      {inputError && <p className="text-red-500 text-sm">{inputError}</p>}
    </section>
  );
}
