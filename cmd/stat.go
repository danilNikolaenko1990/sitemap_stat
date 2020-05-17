package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
	"sitemap_stat/console_report"
)

var (
	reportFileName string
	sitemapUrl     string
	workers        int
)

const (
	SitemapUrlFlagName    = "sitemap-url"
	WorkersFlagName       = "workers"
	OutputFlagName        = "output"
	DefaultReportFilename = "report.csv"
)

func init() {
	defaultWorkers := runtime.NumCPU()
	statCmd.Flags().StringVarP(&sitemapUrl, SitemapUrlFlagName, "u", "", "Sitemap url to fetch sitemap from")
	statCmd.MarkFlagRequired(SitemapUrlFlagName)

	statCmd.Flags().StringVarP(
		&reportFileName,
		"output",
		"o",
		DefaultReportFilename,
		fmt.Sprintf("Report filename, %s by default", DefaultReportFilename))

	statCmd.Flags().IntVarP(&workers, "workers", "w", defaultWorkers, "number of workers, num of CPUs by default")
	rootCmd.AddCommand(statCmd)
}

var statCmd = &cobra.Command{
	Use:   "stat [flags]",
	Short: "Crawl through sitemap and perform requests for each page on sitemap.xml",
	Long:  ``,
	Example: fmt.Sprintf(
		`stat --%s=https://www.google.com/sitemap.xml --%s=report.csv --%s=10`,
		SitemapUrlFlagName,
		OutputFlagName,
		WorkersFlagName),
	Run: func(cmd *cobra.Command, args []string) {
		console_report.Report(sitemapUrl, reportFileName, workers)
	},
}
