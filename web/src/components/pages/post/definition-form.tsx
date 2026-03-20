import { Link, useNavigate } from "@tanstack/react-router";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { definitionSchema, PARTS_OF_SPEECH } from "@/types";
import { Separator } from "@/components/ui/separator";
import { Button } from "@/components/ui/button";
import { Controller, useForm } from "react-hook-form";
import z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { Field, FieldError, FieldLabel } from "@/components/ui/field";
import { Textarea } from "@/components/ui/textarea";
import { capitalize } from "@/lib/utils";

export default function DefinitionForm() {
  const navigate = useNavigate();
  const form = useForm<z.infer<typeof definitionSchema>>({
    resolver: zodResolver(definitionSchema),
    defaultValues: {
      word: "",
      pronunciation: "",
      definition: "",
      partOfSpeech: PARTS_OF_SPEECH[0],
      synonyms: [],
      antonyms: [],
      example: "",
    },
  });

  function handleSubmit(data: z.infer<typeof definitionSchema>) {
    navigate({
      to: "/app/post",
      search: data,
    });
  }

  return (
    <div className="flex flex-col gap-4 p-4 h-full">
      <form
        className="flex flex-col gap-2"
        onSubmit={form.handleSubmit(handleSubmit)}
      >
        <div className="gap-2 grid grid-cols-[2fr_2fr_1fr]">
          <Controller
            name="word"
            control={form.control}
            render={({ field, fieldState }) => (
              <Field data-invalid={fieldState.invalid}>
                <FieldLabel htmlFor={field.name}>Word</FieldLabel>
                <Input
                  {...field}
                  id={field.name}
                  aria-invalid={fieldState.invalid}
                />
                {fieldState.invalid && (
                  <FieldError errors={[fieldState.error]} />
                )}
              </Field>
            )}
          />

          <Controller
            name="pronunciation"
            control={form.control}
            render={({ field, fieldState }) => (
              <Field data-invalid={fieldState.invalid}>
                <FieldLabel htmlFor={field.name}>Pronunciation</FieldLabel>
                <Input
                  {...field}
                  id={field.name}
                  aria-invalid={fieldState.invalid}
                />
                {fieldState.invalid && (
                  <FieldError errors={[fieldState.error]} />
                )}
              </Field>
            )}
          />

          <Controller
            name="partOfSpeech"
            control={form.control}
            render={({ field, fieldState }) => (
              <Field data-invalid={fieldState.invalid}>
                <FieldLabel htmlFor={field.name}>Part Of Speech</FieldLabel>
                <Select>
                  <SelectTrigger>
                    <SelectValue
                      id={field.name}
                      aria-invalid={fieldState.invalid}
                      {...field}
                    />
                  </SelectTrigger>
                  <SelectContent>
                    {PARTS_OF_SPEECH.map((partOfSpeech) => (
                      <SelectItem key={partOfSpeech} value={partOfSpeech}>
                        {capitalize(partOfSpeech)}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
                {fieldState.invalid && (
                  <FieldError errors={[fieldState.error]} />
                )}
              </Field>
            )}
          />
        </div>

        <Controller
          name="definition"
          control={form.control}
          render={({ field, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <FieldLabel htmlFor={field.name}>Definition</FieldLabel>
              <Textarea
                {...field}
                id={field.name}
                aria-invalid={fieldState.invalid}
              />
              {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
            </Field>
          )}
        />

        <Controller
          name="example"
          control={form.control}
          render={({ field, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <FieldLabel htmlFor={field.name}>Example</FieldLabel>
              <Textarea
                {...field}
                id={field.name}
                aria-invalid={fieldState.invalid}
              />
              {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
            </Field>
          )}
        />

        <div className="space-y-2">
          <Label>Synonyms</Label>
          <Input />
        </div>

        <div className="space-y-2">
          <Label>Antonyms</Label>
          <Input />
        </div>

        <Button type="submit" className="mt-2">
          Continue
        </Button>
      </form>
      <Separator>
        <span className="px-4 text-muted-foreground text-xs">OR</span>
      </Separator>
      <Button asChild variant="outline">
        <Link to="/app/define">Define a Word</Link>
      </Button>
    </div>
  );
}
