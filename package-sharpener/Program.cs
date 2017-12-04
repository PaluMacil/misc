using System;
using System.IO;
using System.Linq;

namespace package_sharpener
{
    class Program
    {
        static void Main(string[] args)
        {
            // ensure commands are lowercase
            args = args.Select(a => a.ToLower()).ToArray();

            // exit if no commands
            if (args.Length == 0)
            {
                Console.WriteLine("No actions requested.");
                return;
            }

            if (args.Contains("dedupe"))
            {
                // Get paths; exit if a required folder is missing
                var currentPath = Environment.CurrentDirectory;
                var packagesPath = Path.Combine(currentPath, "packages");
                if (!Directory.Exists(packagesPath))
                {
                    Console.WriteLine($"No packages folder found in ${currentPath}");
                    return;
                }
                var binPath = Path.Combine(currentPath, "bin");
                if (!Directory.Exists(binPath))
                {
                    Console.WriteLine($"No bin folder found in ${currentPath}");
                    return;
                }

                // Get files (with no path) to compare folders
                // TODO: Packages isn't arranged in the same way. Instead of raw files, I need to refactor this to
                // look at folders with a version number in the name and a lib/net40 or lib/net40-client path inside.
                var packageFiles = Directory.EnumerateFiles(packagesPath).Select(f => Path.GetFileName(f));
                var binFiles = Directory.EnumerateFiles(binPath).Select(f => Path.GetFileName(f));

                var dupes = packageFiles.Intersect(binFiles);
                Console.WriteLine($"Duplicated in packages and bin:\n${String.Join(';',dupes)}");
                Console.WriteLine("Delete duplicates from bin?");
                if (PromptYes())
                {
                    string deletePath;
                    foreach (string f in dupes)
                    {
                        deletePath = Path.Combine(currentPath, "bin", f);
                        try
                        {
                            File.Delete(f);
                            Console.WriteLine($"Deleted ${f}");
                        }
                        catch (IOException ioEx)
                        {
                            Console.WriteLine($"Unable to delete ${f}: ${ioEx.Message}");
                        }
                    }
                }
            }
        }

        private static bool PromptYes() { // Is the next typed char a 'Y' or 'y'?
            return Char.ToLower(Console.ReadKey().KeyChar) == 'y';
        }
    }
}
