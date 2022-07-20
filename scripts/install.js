// Reconstruct the optional package name with os.arch() and os.platform()
// Then find it, try to copy it in the main package's /bin/esbuild${extension}
// Then change the bin link in the main package to /bin/esbuild${extension}
// Finally test it with -h (for now)
